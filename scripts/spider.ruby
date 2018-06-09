require 'optparse'
require 'net/http'
require 'net/https'
require 'json'
require 'nokogiri'
require 'open-uri'
require 'watir'
require "awesome_print"

# This will hold the options we parse
options = {}

def trim(text)
  return text.strip.split.join(" ")
end

OptionParser.new do |parser|

  # Whenever we see -n or --url, with an
  # argument, save the argument.
  parser.on("-n", "--url NAME", "The bundle Id of project.") do |v|
    options[:url] = v
  end
end.parse!

# Now we can use the options hash however we like.
if options[:url]
  url = options[:url]
  puts url

  #
  browser = Watir::Browser.new :chrome
  browser.goto url
  browser.link(text: 'View all sizes').click

  # Data
  payload = {}

  # Fetch and parse HTML document
  doc = Nokogiri::HTML(browser.html)

  article = doc.css('.wiki-layout-artist-info article')
  name = trim(article.css('h3').first.content)
  payload['name'] = name

  # Photo info
  article.css('ul li').each do |link|
    link.css('s').each do |tag|
      key = trim(tag.content).delete(':').downcase
      value = link.css('span').first

      if !value.nil?
        payload[key] = trim(value.content)
      end

      if key == 'Dimensions'
        payload[key] = trim(link.content)
      end

    end
  end

  info_tab = doc.css('.wiki-layout-artist-info-wrapper .wiki-layout-artist-info-tab p').first
  if !info_tab.nil?
    payload['info'] = trim info_tab.content
  end

  container = doc.css('div.view-thumnails-sizes-item').first
  linkNode = container.css('div.thumbnail-item a.hidden-blank.ng-binding').last
  payload['url'] = linkNode['href'].partition('!').first
  payload['size'] = linkNode.content

  ap payload

  # Link
  link = article.css('h5 span a').first['href']
  ap link

  # #############################
  # Move to author page
  if !link.nil?

    browser.goto('https://www.wikiart.org' + link)
    doc = Nokogiri::HTML(browser.html)
    author_payload = {}
    name_author = trim(doc.css('.wiki-layout-artist-info article h3').first.content)
    author_payload['name'] = name_author

    article = doc.css('.wiki-layout-artist-info article')
    name = trim(article.css('h3').first.content)
    author_payload['name'] = name

    article.css('ul li').each do |link|
      link.css('s').each do |tag|
        key = trim(tag.content).delete(':').downcase
        value = link.css('span').first

        if !value.nil?
          author_payload[key] = trim(value.content)
        end

        if key == 'Dimensions'
          author_payload[key] = trim(link.content)
        end
      end
    end

    #
    ap author_payload
  end
end

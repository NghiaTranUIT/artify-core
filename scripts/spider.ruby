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
  return text.strip
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

  # Data
  payload = {}

  puts url

  browser = Watir::Browser.new :chrome

  browser.goto 'https://www.wikiart.org/en/caspar-david-friedrich/the-wanderer-above-the-sea-of-fog'
  browser.link(text: 'View all sizes').click

  # Fetch and parse HTML document
  doc = Nokogiri::HTML(browser.html)

  name = trim(doc.css('.wiki-layout-artist-info article h3').first.content)
  payload['name'] = name

  doc.css('.wiki-layout-artist-info article ul li').each do |link|
    link.css('s').each do |tag|
      key = trim(tag.content).delete(':').downcase
      value = link.css('span')[0]

      if !value.nil?
        payload[key] = trim(value.content)
      end

      if key == 'Dimensions'
        payload[key] = trim(link.content)
      end

    end
  end

  container = doc.css('div.view-thumnails-sizes-item').first
  linkNode = container.css('div.thumbnail-item a.hidden-blank.ng-binding').last
  payload['url'] = linkNode['href'].partition('!')[0]
  payload['size'] = linkNode.content

  ap payload
end

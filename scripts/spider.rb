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

  puts "Fetching ... " + url

  #
  browser = Watir::Browser.new :chrome
  browser.goto url
  browser.link(text: 'View all sizes').click

  # Data
  payload = {}
  payload['original_source'] = url

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

  # Get size
  size = payload['size']
  if !size.nil?
    components = size.split("x")
    payload['width'] = components.first.to_i
    payload['height'] = components.last.to_i
  end

  # Link
  link = article.css('h5 span a').first['href']
  ap "Go to ... " + link

  # #############################
  # Move to author page
  if !link.nil?
    author_url = 'https://www.wikiart.org' + link
    browser.goto(author_url)
    doc = Nokogiri::HTML(browser.html)
    author_payload = {}
    author_payload['original_source'] = author_url

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

        if key == "Dimensions"
          author_payload[key] = trim(link.content)
        end
      end
    end
  end

  # Full payload
  body = {
    "photo" => payload,
    "author" => author_payload
  }
  ap body

  endpoint = "https://api.proxyman.app/spider/wikiart"
  ap "Posting to #{endpoint}"

  # Request
  encoded_url = URI.encode(endpoint)
  uri = URI.parse(encoded_url)
  https = Net::HTTP.new(uri.host, 443)
  https.use_ssl = true
  request = Net::HTTP::Post.new(uri.path, 'Content-Type' => 'application/json')
  request.body = body.to_json
  response = https.request(request)

  ap response
  ap "code #{response.code}"

  ## Succss
  if response.code.to_i == 200
    ap "Success !!!"
    value = JSON.parse(response.body)
    puts value
  else
    ap "ERROR!!! at " + url
  end

end

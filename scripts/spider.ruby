require 'optparse'
require 'net/http'
require 'net/https'
require 'json'
require 'nokogiri'
require 'open-uri'
require 'watir'

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

  puts url

  browser = Watir::Browser.new :chrome

  browser.goto 'https://www.wikiart.org/en/caspar-david-friedrich/the-wanderer-above-the-sea-of-fog'
  browser.link(text: 'View all sizes').click

  # Fetch and parse HTML document
  doc = Nokogiri::HTML(browser.html)

  name = trim(doc.css('.wiki-layout-artist-info article h3').first.content)
  puts name

  doc.css('.wiki-layout-artist-info article ul li').each do |link|
    link.css('s').each do |tag|
      key = trim(tag.content).delete(':')
      value = link.css('span')[0]
      puts key

      if !value.nil?
        puts trim(value.content)
      end

      if key == 'Dimensions'
        puts trim(link.content)
      end

    end
    # puts link.content
  end

  puts "### Click Btn"
  container = doc.css('div.view-thumnails-sizes-item').first
  container.css('div.thumbnail-item a.hidden-blank.ng-binding').each do |link|
    puts link.content + " = " + link['href'].partition('!')[0]
  end

end

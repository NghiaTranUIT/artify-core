require 'optparse'
require 'net/http'
require 'net/https'
require "awesome_print"
require 'csv'

# This will hold the options we parse
options = {}

OptionParser.new do |parser|

  # Whenever we see -n or --url, with an
  # argument, save the argument.
  parser.on("-n", "--name NAME", "The CSV name.") do |v|
    options[:name] = v
  end
end.parse!

# Now we can use the options hash however we like.
if options[:name]
  name = options[:name]

  puts "Reading CSV " + name + " ..."
  csv_text = File.read(name)
  csv = CSV.parse(csv_text, :headers => true)
  csv.each do |row|
    url = row["url"]
    script = 'ruby spider.ruby --url ' + url
    system script
  end
end

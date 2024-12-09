# frozen_string_literal: true

require 'net/http'
require 'uri'
require 'rack'

app = proc do |env|
  req = Rack::Request.new(env)

  sleep = req.params['sleep']

  target_url = URI("http://rust:3000/?sleep=#{URI.encode_www_form_component(sleep)}")

  begin
    response = Net::HTTP.get_response(target_url)

    next [200, { 'Content-Type' => 'text/plain' }, [response.body]] if response.is_a?(Net::HTTPSuccess)
  rescue
    # noop
  end

  [500, { 'Content-Type' => 'text/plain' }, ['Failed to fetch data']]
end

run app

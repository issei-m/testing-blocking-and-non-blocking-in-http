# frozen_string_literal: true

port ENV.fetch('PORT') { 8000 }

threads_count = ENV.fetch('RAILS_MAX_THREADS') { 5 }.to_i
threads threads_count, threads_count

worker_timeout 3600

environment 'production'

log_requests

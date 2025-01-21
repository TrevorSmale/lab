#!/usr/bin/perl
use strict;
use warnings;
use LWP::Simple;
use JSON;

# Define the URL to fetch JSON data
my $url = 'https://jsonplaceholder.typicode.com/posts/1';

# Fetch the JSON data
my $json_data = get($url);

# Check if the fetch was successful
die "Failed to fetch data from $url\n" unless defined $json_data;

# Decode JSON data into a Perl hash
my $decoded_json = decode_json($json_data);

# Print the parsed data
print "Title: $decoded_json->{title}\n";
print "Body: $decoded_json->{body}\n";
package com.jeffotoni.sclient.controller;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(path = "/v1/client")
public class ClientController {

	private static final String SECONDS = null;

	@GetMapping
	//@ResponseBody
	@ResponseStatus(HttpStatus.OK)
	public ResponseEntity<String> client() throws URISyntaxException, IOException, InterruptedException {

		HttpClient httpClient = HttpClient.newBuilder()
		.version(HttpClient.Version.HTTP_1_1)
		.build();
		
		HttpRequest request = HttpRequest.newBuilder()
		.uri(new URI("http://localhost:3000/v1/customer"))
		.GET()
		.setHeader("Content-Type", "application/json")
		.setHeader("User-Agent", "Java 11 HttpClient Bot")
		.build();
		
		HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

		return ResponseEntity.created(URI.create(String.format("/v1/client"))).
				header("Engine", "Spring Boot")
				.header("Content-Type", "application/json")
				.body(response.body());
	}
}

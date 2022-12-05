package com.jeffotoni.sclient.controller;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(path = "/v1/client")
public class ClientController {
	@Value("#{environment.DOMAIN}")
	private String DOMAIN = "http://127.0.0.1:3000/v1/client";

	@GetMapping
	@ResponseStatus(HttpStatus.OK)
	public ResponseEntity<String> client() throws URISyntaxException, IOException, InterruptedException {
		HttpRequest request = null;
		try {

			// if (DOMAIN.length() == 0 ) {
			// 	DOMAIN = "http://127.0.0.1:3000/v1/client";
			// }
			
			System.out.println(DOMAIN);

			HttpClient httpClient = HttpClient.newBuilder()
			.version(HttpClient.Version.HTTP_1_1)
			.build();
			
			request = HttpRequest.newBuilder()
			.uri(new URI(DOMAIN))
			.GET()
			.setHeader("Content-Type", "application/json")
			.setHeader("User-Agent", "Java 11 HttpClient Bot")
			.build();
		
			HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

			return ResponseEntity.created(URI.create(String.format("/v1/client"))).
				header("Engine", "Spring Boot")
				.header("Content-Type", "application/json")
				.body(response.body());

		} catch(URISyntaxException | IOException | InterruptedException e) {
			System.err.println(e.toString());

			if (request == null) {
				return null;
			}
			return ResponseEntity.created(URI.create(String.format("/v1/client"))).
				header("Engine", "Spring Boot")
				.header("Content-Type", "application/json")
				.body(e.getMessage());
		}
	}
}

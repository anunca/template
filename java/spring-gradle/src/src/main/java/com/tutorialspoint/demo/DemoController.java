package com.tutorialspoint.demo;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.http.MediaType;

@RestController
public class DemoController {

	@RequestMapping(value = "/")
	public String home() {
		return "Home";
	}

	@RequestMapping(value = "/hello")
	public String hello() {
		return "Hello";
	}

	@GetMapping(value="/index", produces = MediaType.TEXT_PLAIN_VALUE)
	public String index() {
			return "Index";
	}
}

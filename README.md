# Simple URL Shortner 

By passing a JSON object with a URL and shortened URL, the program adds these to a map indexed by a short URL. The long URL can then be queried by passing a short URL to the API.

Example JSON request to get a URL. API route: `/get`
```JSON
{
    "short": "x78d0n"
}
```

Example response
```JSON
{
    "google.com"
}
```

Example JSON request to add a short URL. API route: `/add`
```JSON
{
	"short": "x78d0n",
	"long": "google.com"
}
```

Example JSON response to add a short URL. API route: `/add`
```JSON
{
    "x78d0n": "google.com",
}
```



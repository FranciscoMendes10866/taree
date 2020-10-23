# Project idea

The idea of creating this project is due to the fact that you can explore the world of Nginx more and learn more about how to use your reverse proxy and become more familiar with this new technology (for me).

## How to use

First step (build the images):

```bash
docker-compose build
```

Second step (orchestrate and create the containers):

```bash
docker-compose up -d
```

Third step (visit the website):

```
http://localhost
```

## Nice to try

To get the JSON response from the `api-01`, visit the following URL:

```
http://localhost/first-api
```

To get the JSON response from the `api-02`, visit the following URL:

```
http://localhost/second-api
```

## Summary

Through the URL that is visited by the user, the nginx reverse proxy will redirect the corresponding service.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
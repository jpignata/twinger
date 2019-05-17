# Twinger

The application is a daemon that speaks a subset of the [finger user information protocol][0]
and returns information about Twitter users.

![](doc/screenshot.png)

# Usage

1. **[Get a Twitter API Key][1]**

1. **Clone the repo**

    ```console
    git clone https://github.com/jpignata/twinger.git
    ```

1. **Build the project**

    ```console
    make
    ```

1. **Run it**

    ```console
      CONSUMER_KEY=yourconsumerkey \
      CONSUMER_SECRET=yourconsumersecret \
      ACCESS_TOKEN=youraccesstoken \
      ACCESS_TOKEN_SECRET=youraccesstokensecret sudo ./twinger
    ```

1. **lol**

    ```console
    $ finger washingtonpost@localhost
    [localhost]
    Trying 127.0.0.1...
    Login: washingtonpost                    Name: The Washington Post
    Location: Washington, DC                 URL: http://t.co/Hq7hTYkOPg
    Description: Breaking news, analysis, and opinion. Founded in 1877. Our staff on Twitter: https://t.co/VV0UBAMHg8
    Last tweet Fri May 17 17:45:36 +0000 2019 from SocialFlow
    Plan: Analysis: Florida lawmakers rail against FBI for secrecy on voter breaches https://t.co/EgWZFFXmyu
    ```

[0]: https://tools.ietf.org/html/rfc1288
[1]: https://developer.twitter.com

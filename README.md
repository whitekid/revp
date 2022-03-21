# revp; reverse proxy for developers

## Usage

in your laptop, start very sample web server

    $ python -m http.server
    Serving HTTP on :: port 8000 (http://[::]:8000/) ...

open new terminal and start `revp` client with demo secret

    $ revp --secret demo 127.0.0.1:8000
    forwarding http://revp.woosum.net:57943/ -> 127.0.0.1:8000

open the forwarded URL.

## allow firewall

    # for server connection
    ufw allow 49999/tcp

    # for proxy connection
    ufw allow 50000:59999/tcp

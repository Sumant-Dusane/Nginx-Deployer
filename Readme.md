# NGINX DEPLOYER

## Installation
`curl -fsSL https://raw.githubusercontent.com/Sumant-Dusane/nginx-deployer/master/install.sh`

## MVP
1. Scans the folders
2. Takes input for deployment
3. Takes input for domain name or ip address
4. Show loader while config file is written
5. Show success for http dep and then ask whether to deploy https via certbot
6. Run certbot command and restart nginx
7. Show success message with domain name or ip addr


## Future Scopes
- Cross OS versions mac, win, linux
- Security & Headers
- Load Balancer
- Rate Limiting
- Caching and static files
- Gzip Compression
# Install DDNS

## Install duck dns

1. Get your token in the Duck DNS website;
2. Change the INTERFACE and TOKEN according to your network configuration;
3. Copy the folder `./duck_dns` to the folder `/opt` of the host server
```
sudo cp -r ./duckdns /opt
```
4. Append the line bellow in the root crontab:
```
sudo crontab -e
# Add the line '*/5 * * * * /opt/duck_dns/duck.sh >/dev/null 2>&1'
```

# Device discovery for Philips Hue
From: https://developers.meethue.com/develop/application-design-guidance/hue-bridge-discovery/#1.%20UPnP

## Tools
- sudo apt install gupnp-tools for gnu-upnp tools
- ip link

1. Find the interfaces to scan for devices
> ip link up

2. SSDP search for devices
> gssdp-discover -i <interface>

3. grep for Hue device
> ...|grep -i ipbridge

4. If this above doesn't return any device use n-pnp
> curl -i https://discovery.meethue.com/

You can also use nmap to check the available ports of the Hue bridge
> nmap -A -T4 <ipadddress>

This will show that there are three servers running; HTTP:80, HTTPS:443, DMTF-CIM:8080
There is a basic description.xml endpoint for metadata about the bridge
> curl -i <ipaddress>/description.xml

The cert seem to be self signed, so use -k to allow insecure
> curl -k -i https://<ipaddress>/description.xml

TODO(dape): how do we integrate with the https endpoint without -k
TODO(dape): on a real deployment the bridge should be hidden to prevent https acess
TODO(dape): on a real deployment the curl -i discovery should be prevent by guests (or all)
TODO(dape): investigate DMTF-CIM; what can this be used for? Seem to be undocumented
TODO(dape): check if Hue app for Android/iOS uses HTTPS? (otherwise it is very very simple to hack)
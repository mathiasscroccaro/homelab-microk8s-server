INTERFACE="wlp2s0"
TOKEN=""

IPV6=$(ip -6 addr show dev $INTERFACE scope global | sed -e's/^.*inet6 \([^ ]*\)\/.*$/\1/;t;d')

echo url="https://www.duckdns.org/update?domains=mathias-dev&token=${TOKEN}&ipv6=${IPV6}" | curl -k -o /opt/duck_dns/duck.log -K -

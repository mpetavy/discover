#/bin/sh
tar -zxvf discover-1.0.0-1234-linux-amd64.tar.gz -C /tmp ./discover
lxc delete discover --force
lxc launch images:debian/10 discover
lxc file push /tmp/discover discover/root/
lxc exec discover -- /root/discover -log.verbose -log.file=/root/discover.log -cfg.file=/root/discover.json -app.product=edgebox -callback.noerrors -service=install
lxc exec discover -- systemctl enable discover.service
lxc exec discover -- systemctl start discover.service
# lxc exec discover -- /bin/bash
lxc export discover discover-1.0.0-1234-lxc.tar.gz --instance-only

# ! bin/sh
#
while
fail2ban-client status sshd > fail.txt
last > ssh.txt
netstat -s > net.txt
sleep 30

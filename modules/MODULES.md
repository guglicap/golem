# cpu
Displays information about the current CPU usage.

### Options

* **Format**


   * "%core"
 
      Core name. Example: "cpu0"

   * "%usage" 

      Usage in percent. Example: "5%"

 
  * **Default is "%core: %usage"**

* **PerCPU**  
   When`true`displays statistics for every cpu. On a quad core machine, for example, the output will contain cpu0, cpu1, cpu2, cpu3.  
   **Default is true.**

# date
Displays the current date/time.

### Options

   * **Format**
     
   This program uses Go constants for time format. For reference look [here](https://golang.org/src/time/format.go?s=14285:14327#L83). Some examples:

      *  "15:04" hour:minute
      *  "02/01/2006" dd/mm/yyyy

  **Default is "02/01/2006 15:04"**


# disk

Displays info about disk usage.

### Options 

   * **Mountpoint**  
      Set the mountpoint of the filesystem.  
      **Default is `/`**

   * **Format**  
   
     * "%size": Disk size in GB. 

     * "%used": Usage in GB.

     * "%avail": Free space in GB.

     * "%usePerc": Percentage of usage.

     * "%mount": The mountpoint.
     
     **Default is "%mount: %usedGB/%sizeGB"**  



# launcher

Displays an icon tray.

### Options 

   * **Programs**       
      Comma separated list of commands.  
      **Default is "firefox"**

   * **Icons**       
    Comma separated list of icons.  
    These actually need not be icons, they can be text, but using FontAwesome icons is recommended.   
    An underscore uses the program name as icon.  
    **Default is "\uf269"**

# mem

Displays information about current Virtual Memory usage.  

### Options

   * **Format**
      
       * "%total": Total amount of memory in MBs.
       
       * "%free":  Free memory in MBs.

       * "%avail": Available memory in MBs.
 
       * "%used":  Used memory in MBs.

       * "%usePerc": Percentage of memory used.

  **Default is %usedMB / %totalMB**


# text

Displays a text and returns. It's also useful to pad the bar content, have a look at the example config file.

### Options 

   * **Text**  
 The text to display.

# button

Displays a button.

### Options
     
   * **Command**   
     The command to run on button pressed.

   * **Text**  
     Text to be used as button. It can, of course, be a FontAwesome icon. 

# mpd

Displays mpd related info and controls. You also need to have`mpc`installed in order for this to work.

### Options

   * **Format**
     
     * "%song": Info about the current song, formatted according to SongFormat

     * "%toggle": Toggle button.  
    
     * "%prev": Previous song button.

     * "%next": Next song button.

  **Default is "%song %toggle"**

   * **SongFormat**
    
     * "%title":  Song title.

     * "%artist": Song artist.

     * "%album":  Song album.

  **Default is "%artist - %title"**

   * **PrevButton**  
 Icon for the previous song button  
**Default is "\uf04a"**

   * **NextButton**  
 Icon for the next song button  
     **Default is "\uf04e"**

   * **TogglePlaying**  
 Icon for the toggle button when playing  
**Default is "\uf04c"**

   * **TogglePaused**:  
Icon for the toggle button when paused  
**Default is "\uf04b"**

   * **Address**  
mpd server address  
**Default is ":6600"**

   * **Password**  
 mpd server password  
**Default is ""**

# net

Displays the machine IP address

### Options

   * **Interface**  
The network interface.  
**Default is "enp3s0"**

# power

Displays a powertray.

### Options

   * **Format**
   
     * "%P": Poweroff button.
    
     * "%R": Reboot button.
  
     * "%S": Suspend button.

  **Default is "%P"**

   * **PowerOffText**  
Text to use for the poweroff button.  
 Can be an icon.  
**Default is "\uf011"**

   * **RebootText**  
Text to use for the reboot button.  
Can be an icon.  
**Default is "\uf021"**

   * **SuspendText**  
Text to use for the suspend button.  
Can be an icon.  
**Default is "\uf186"**

   * **PowerOffCmd**  
Command to power the pc off.  
**Default is "poweroff"**

   * **RebootCmd**  
Command to reboot the pc.  
**Default is "reboot"**

   * **SuspendCmd**  
Command to suspend the pc.  
**Default is "systemctl suspend"**

# whoami

Displays info about the current user

### Options

   * **Format**  
  
      * "%uname": Login name.

      * "%name":  
 User first name. On my system it's an empty string. See [here](https://golang.org/pkg/os/user/#User).

      * "%uid": UID of the current user.

      * "%gid": GID of the current user.
   
   **Default is "%uname"**


# ws

Displays a workspace indicator, it's currently only working for bspwm.

### Options

   * **WsFocused**  
Icon to use for the active desktop.  
**Default is "\uf111"** 


   * **WsUnfocused**  
Icon to use for the inactive desktops.  
**Default is "\uf10c"** 
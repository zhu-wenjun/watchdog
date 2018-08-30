# watchdog

1. The system will reboot once the "/dev/watchdog" device is
   opened and no subsequent data writes to the device in 3 minutes
   (It depends on your watchdog driver setting).

2. The watchdog daemon writes a data to "/dev/watchdog" once per minute.

3. Use "echo V > /dev/watchdog" to stop the watchdog and keep the system
   from rebooting.

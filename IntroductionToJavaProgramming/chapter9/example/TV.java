package chapter9.example;


public class TV {
    int channel = 1;
    int volumelevel = 1;
    boolean on = false;

    public TV() {    }

    public void turnOn() {
    on = true;}

    public void turnOff() {
    on = false;}

    public void setChannel (int newChannel) {
        if (on && newChannel >= 1 && newChannel <= 120)
            channel = newChannel;
    }

    public void setVolume (int newVolumelevel) {
        if (on && newVolumelevel >= 1 && newVolumelevel <=7)
            volumelevel = newVolumelevel;
    }

    public void channelUp() {
        if (on && channel <120)
            channel++;
    }

    public void channelDown() {
        if (on && channel > 1)
            channel--;
    }

    public void volumeUp() {
        if (on && volumelevel < 7)
            volumelevel++;
    }

    public void volumeDowm() {
        if (on && volumelevel > 1)
            volumelevel--;
    }
}
# Basement of AndroidStudio
+ adb shell命令可以用来调用一系列手机资源
    + 把文件传到电脑上去
    ```shell
    adb pull fileName destAddrInPC
    ```


---
+ 存文件(外部)
    ```java
    String time = new String(hour+":"+minute+":"+second+" "+mSecond+" ");
    File sdCard = Environment.getExternalStorageDirectory();
    // 获取的sd卡的绝对路径为 /storage/sdcard
    sdCard = new File(sdCard, "/AAirScreen");
    sdCard.mkdirs();// 创建MyFiles目录
    sdCard = new File(sdCard, fileName);
    FileOutputStream out = new FileOutputStream(sdCard,true);
    Writer writer = new OutputStreamWriter(out);
    try {
        writer.write(time+content);
    } finally {
        writer.close();
    }
    ```
+ 删除文件(外部)
    ```java
    public static void deletePath() {
        File file;
        file = Environment.getExternalStorageDirectory();
        file = new File(file.getPath() + "/AAirScreen");
        RecursionDeleteFile(file);
    }
    public static void RecursionDeleteFile(File file) {
        if (file.isFile()) {
            file.delete();
            return;
        }
        if (file.isDirectory()) {
            File[] childFile = file.listFiles();
            if (childFile == null || childFile.length == 0) {
                file.delete();
                return;
            }
            for (File f : childFile) {
                RecursionDeleteFile(f);
            }
            file.delete();
        }
    }
    ```
#ifdef __ANDROID__
    #include <android/log.h>

    // Redirect stderr to stdout
    #define stderr stdout

    // Define a logging macro that can be used explicitly
    #define SE_LOG(...) __android_log_print(ANDROID_LOG_ERROR, "SwissEph", __VA_ARGS__)
#endif
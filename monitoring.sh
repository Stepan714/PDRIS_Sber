#!/bin/bash

PID_FILE="/tmp/monitor.pid"
LOG_DIR="$HOME/Documents/psid_sber/lab1"

monitor() {
    while true; do
        timestamp=$(date +"%Y-%m-%d_%H-%M-%S")
        date_for_filename=$(date +"%Y-%m-%d")
        log_file="${LOG_DIR}/monitor_${timestamp}_${date_for_filename}.csv"

        df_output=$(df -h)
        inode_output=$(df -i)

        echo "Timestamp, Disk Usage, Inode Usage" > "$log_file"
        echo "$timestamp, $df_output, $inode_output" >> "$log_file"

        sleep 86400 # 24 часа
    done
}

case "$1" in
    START)
        if [ -f "$PID_FILE" ]; then
            echo "Мониторинг уже запущен с PID $(cat $PID_FILE)"
            exit 1
        fi
        monitor &
        echo $! > "$PID_FILE"
        echo "Мониторинг запущен с PID $!"
        ;;
    
    STOP)
        if [ -f "$PID_FILE" ]; then
            kill $(cat "$PID_FILE") && rm -f "$PID_FILE"
            echo "Мониторинг остановлен"
        else
            echo "Мониторинг не запущен"
        fi
        ;;
    
    STATUS)
        if [ -f "$PID_FILE" ]; then
            echo "Мониторинг запущен с PID $(cat $PID_FILE)"
        else
            echo "Мониторинг не запущен"
        fi
        ;;
    
    *)
        echo "Использование: $0 {START|STOP|STATUS}"
        exit 1
        ;;
esac


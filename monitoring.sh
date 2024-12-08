#!/bin/bash

PID_FILE="$HOME/monitor.pid"
LOG_DIR="monitorings"
mkdir -p "$LOG_DIR"

monitor() {
    current_date=$(date +'%Y-%m-%d')
    log_file="$LOG_DIR/monitoring_${current_date}.csv"

    if [ ! -f "$log_file" ]; then
        echo "Timestamp,Disk Usage (%),Free Inodes" > "$log_file"
    fi

    while true; do
        timestamp=$(date '+%Y-%m-%d %H:%M:%S')
        new_date=$(date +'%Y-%m-%d')

        if [ "$new_date" != "$current_date" ]; then
            current_date=$new_date
            log_file="$LOG_DIR/monitoring_${current_date}.csv"
            echo "Timestamp,Disk Usage (%),Free Inodes" > "$log_file"
        fi

        disk_usage=$(df / | tail -1 | awk '{print $5}' | sed 's/%//')
        free_inodes=$(df -i / | tail -1 | awk '{print $4}')

        echo "$timestamp,$disk_usage,$free_inodes" >> "$log_file"
        
        sleep 20
    done
}

case "$1" in
    START)
        if [ -f "$PID_FILE" ]; then
            echo "Мониторинг уже запущен с PID $(cat "$PID_FILE")"
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
            echo "Мониторинг запущен с PID $(cat "$PID_FILE")"
        else
            echo "Мониторинг не запущен"
        fi
        ;;
    
    *)
        echo "Использование: $0 {START|STOP|STATUS}"
        exit 1
        ;;
esac

#ifndef _SERIAL_H_
#define _SERIAL_H_

#define SERIAL_COM1_BASE                0x3F8
#define SERIAL_COM2_BASE                0x2F8

    /*
    名稱	        Offset	說明
    Data Register	+0	    寫入/讀取資料
    FIFO Control	+2	    控制 FIFO 緩衝
    Line Control	+3	    設定資料位元、停止位元、Parity，含 DLAB
    Modem Control	+4	    控制 RTS/CTS 等 modem 訊號
    Line Status	    +5	    狀態讀取，例如 Transmit FIFO 空等
    */
#define SERIAL_DATA_PORT(base)          (base)
#define SERIAL_FIFO_COMMAND_PORT(base)  (base + 2)
#define SERIAL_LINE_COMMAND_PORT(base)  (base + 3)
#define SERIAL_MODEM_COMMAND_PORT(base) (base + 4)
#define SERIAL_LINE_STATUS_PORT(base)   (base + 5)

#define SERIAL_LINE_ENABLE_DLAB         0x80

void serial_configure_baud_rate(unsigned short com, unsigned short divisor);
void serial_configure_line(unsigned short com);

 /** serial_is_transmit_fifo_empty:
  *  Checks whether the transmit FIFO queue is empty or not for the given COM
  *  port.
  *
  *  @param  com The COM port
  *  @return 0 if the transmit FIFO queue is not empty
  *          1 if the transmit FIFO queue is empty
  */
int serial_is_transmit_fifo_empty(unsigned int com);

void serial_write(char *buf, unsigned int len);

#endif
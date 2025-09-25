# VISCA Serial Command Reference (Sony FCB-EX480AP)

This document lists commonly used VISCA commands and inquiries for the Sony FCB-series
camera blocks (tested on **FCB-EX480AP**). These are transmitted over UART (TTL 5V)
at **9600 8N1** (default baud, can also support 19.2k and 38.4k).

---

## 📡 VISCA Packet Format

```
[Address] [Command Bytes...] FF
```

- **Address:** `0x81` for Camera 1 (default).
- **Command Bytes:** Action or inquiry data.
- **FF:** Terminator.

Camera replies usually start with `0x90`.

---

## ⚡ Basic Commands

| Command                          | Bytes (Hex)                  | Notes |
|----------------------------------|------------------------------|-------|
| **Power On**                     | `81 01 04 00 02 FF`          | Turns camera on |
| **Power Off**                    | `81 01 04 00 03 FF`          | Turns camera off |
| **Zoom Tele (variable speed)**   | `81 01 04 07 2p FF`          | `p=0x0–0x7` (low→high speed) |
| **Zoom Wide (variable speed)**   | `81 01 04 07 3p FF`          | `p=0x0–0x7` |
| **Zoom Stop**                    | `81 01 04 07 00 FF`          | Stops zoom |
| **Focus Auto**                   | `81 01 04 38 02 FF`          | Enable AF |
| **Focus Manual**                 | `81 01 04 38 03 FF`          | Switch to MF |
| **Focus Near (variable)**        | `81 01 04 08 3p FF`          | Move focus nearer |
| **Focus Far (variable)**         | `81 01 04 08 2p FF`          | Move focus farther |
| **Focus Stop**                   | `81 01 04 08 00 FF`          | Stops focus move |
| **Iris Auto**                    | `81 01 04 0B 00 FF`          | Auto iris |
| **Iris Manual Up**               | `81 01 04 0B 02 FF`          | Increase iris |
| **Iris Manual Down**             | `81 01 04 0B 03 FF`          | Decrease iris |

---

## 🔎 Inquiries (status checks)

| Inquiry                          | Bytes (Hex)                  | Example Reply | Notes |
|----------------------------------|------------------------------|---------------|-------|
| **Power Inquiry**                | `81 09 04 00 FF`             | `90 50 0p FF` (`p=2 on, 3 off`) | Check power state |
| **Zoom Position Inquiry**        | `81 09 04 47 FF`             | `90 50 0p 0q 0r 0s FF` | 4 nibbles = zoom pos |
| **Focus Mode Inquiry**           | `81 09 04 38 FF`             | `90 50 0x FF` (`2=Auto, 3=Manual`) | |
| **Focus Position Inquiry**       | `81 09 04 48 FF`             | `90 50 0p 0q 0r 0s FF` | 4 nibbles = focus pos |
| **Iris Inquiry**                 | `81 09 04 0B FF`             | `90 50 0p 0q FF` | Current iris setting |
| **Gain Inquiry**                 | `81 09 04 0C FF`             | `90 50 0p 0q FF` | Gain value |
| **Shutter Inquiry**              | `81 09 04 0A FF`             | `90 50 0p 0q FF` | Shutter speed |
| **WB Mode Inquiry**              | `81 09 04 35 FF`             | `90 50 0p FF` | WB mode (0=Auto, 1=Indoor, etc.) |

---

## ✅ Example: Checking if Camera is Alive

Send:
```
81 09 04 47 FF
```
(Zoom Position Inquiry)

Expected Reply:
```
90 50 0p 0q 0r 0s FF
```
If a valid reply is received, the UART link is working.

---

## 🛠 Notes

- Always terminate packets with `FF`.
- Camera address defaults to `0x81`. Changeable if daisy-chaining multiple cameras.
- Replies begin with `0x90`.
- Timeouts: allow 50–200 ms for reply.
- Error replies start with `90 60 ...`.

---

## 📚 References

- Sony FCB-EX480AP Service Manual (Section 4: Command List)
- Sony VISCA Protocol Specifications

# Requirement: pip install PyMuPDF
# How to run: python document.py -h
#Example Input:
#python document.py -m sign -s sample.pdf -d destination.pdf -i sign.png -x 484.296875 -y 150.40625 -w 128.0 -t 123.0 -p 2

import fitz
from datetime import datetime
import sys, getopt

def signDocument(x0, y0, width, height, page, src_pdf, dst_pdf, src_img):
    x1 = x0 + width
    y1 = y0 + height
    page = page - 1
    img_rect = fitz.Rect(x0, y0, x1, y1) #kiri, atas, kanan, bawah
    document = fitz.open(src_pdf)

    page = document[page]
    page.insert_image(img_rect, filename=src_img)

    document.save(dst_pdf)
    document.close()

def waterMarking(x0, y0, width, height, page, src_pdf, dst_pdf):
    now = datetime.now()
    dt = now.strftime("%d/%m/%Y %H:%M:%S")
    
    x1 = x0 + width
    y1 = y0 + height
    page = page - 1

    text_rect = fitz.Rect(x0, y0, x1, y1) #kiri, atas, kanan, bawah
    document = fitz.open(src_pdf)
    text = "smartsign " + dt
    page = document[page]
    page.insert_textbox(text_rect, text, fontsize =0,
                   fontname = "Times-Roman",      
                   fontfile = None,               
                   align = 1)                     

    document.saveIncr()
    document.close()

def main(argv):
    # mode = ''; x = 10; y = 10
    # width = 5; height = 5; page = 1
    # src_pdf = ''; dst_pdf = ''; src_img = ''
    opts, args = getopt.getopt(argv,"hm:s:d:i:x:y:w:t:p:",["mode=","src_pdf=","dst_pdf=","src_img=","x=","y=","width=","height=","page="])
    for opt, arg in opts:
        if opt == '-help':
            print('RUN WITH document.py -m <mode> -s <src_pdf> -d <dst_pdf> -i <src_img> -x <x> -y <y> -w <width> -t <height> -p <page>')
            sys.exit()
        elif opt in ("-m", "--mode"):
            mode = arg
        elif opt in ("-s", "--src_pdf"):
            src_pdf = arg
        elif opt in ("-d", "--dst_pdf"):
            dst_pdf = arg
        elif opt in ("-i", "--src_img"):
            src_img = arg
        elif opt in ("-x", "--x"):
            x = float(arg)
        elif opt in ("-y", "--y"):
            y = float(arg)
        elif opt in ("-w", "--width"):
            width = float(arg)
        elif opt in ("-t", "--height"):
            height = float(arg)
        elif opt in ("-p", "--page"):
            page = int(arg)
     
    if mode == 'watermark':
        waterMarking(x, y, width, height, page, src_pdf, dst_pdf)
    elif mode == 'sign':
        signDocument(x, y, width, height, page, src_pdf, dst_pdf, src_img)
    else:
        print('Mode not found, Please select sign or watermark')

if __name__ == '__main__':
    main(sys.argv[1:])
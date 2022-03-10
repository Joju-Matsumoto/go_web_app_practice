cd xml/wordnet/
file="jpn_wn_lmf.xml"

if [ ! -e "${file}.gz" ]; then
wget http://compling.hss.ntu.edu.sg/wnja/data/1.1/jpn_wn_lmf.xml.gz
fi

if [ ! -e $file ]; then
gzip -dk "${file}.gz"
fi
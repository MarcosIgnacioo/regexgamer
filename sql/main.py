# ([a-zA-Z]+[\d]*?#?)*\*?<?>?=?((’?'?)*([\d]*[a-zA-Z]*)*(’?'?)*)?
# ([a-za-z]+[\d]*?#?)*\*?<?>?=?((‘?’?'?)*([\d]*[a-za-z]*)*(‘?’?'?)*)?
# from lexer import KEYWORDS

# Lo mas cercano a funcionar justo como ocupamos
# ([a-zA-Z]+[\d]*?#?)*\.?,?\*?<?>?=?((‘?’?'?)*([\d]*[a-zA-Z]*)*(‘?’?'?)*)?
# Esta es la que jala

# ([‘’']*[\d*a-zA-Z ?]*['’])?([a-zA-Z\.?#?\d]*)(,?<?>?=?\*?\(?\)?#?)

# asdf = KEYWORDS.get("SELECT")
# if asdf == None:
#     print("nohay")
# else:
#     print("sihay")
import re

txt = "'adsf dasf'"
x = re.findall("(('|‘|’)*[a-zA-Z]* [a-zA-Z]*('|‘|’)*)", txt)
print(x)

import re
from math import log

def main():
    with open("./books/pride_and_prejudice.txt") as f:
        pnp = f.read()
    with open("./books/frankenstein.txt") as q:
       f = q.read()
    corpus = [pnp, f]
    dt_tf = [
    for i in range(len(corpus)):
        unique_words, all_words = tokenize(corpus[i])
        for w in range(all_words):
            dt_tf[w][i] = dt_tf[w][i] + (1/len(all_words))



def inverse_docuement_frequency(first_list, *argv):
    size_of_corpus = len(argv) + 1
    appearences = {}
    vector_res = {}
    if len(argv) == 0:
        print("Only 1 file")
        return
    for word in first_list:
        if word != appearences:
            appearences[word] = 1
        else:
            appearences[word] += 1
    for arg in argv:
        for word in arg:
            if word != appearences:
                appearences[word] = 1
            else:
                appearences[word] += 1
    for i in appearences:
        if i not in first_list:
            continue
        elif i != "total_words": 
            tf = first_list[i]/first_list["total_words"]
            dt = log(size_of_corpus/appearences[i], 10)
            vector_res[i] = tf * dt
    print(vector_res)



def tokenize(string)->list:
    tokens = []
    all_words = []
    words = string.split()
    for word in words:
        word = word.lower()
        word = re.sub(r'[^\w\s]','',word)
        if word not in tokens:
            tokens.append(word)
        all_words.append(word)
    return tokens, all_words


main()
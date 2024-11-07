import re

def main():
    pnp_words, pnp_freq = generate_frequency_list("./books/pride_and_prejudice.txt", 10)
      
    f_words, frankenstein_freq = generate_frequency_list("./books/frankenstein.txt", 10)
    
    print (pnp_words, len(pnp_freq))
    print(f_words, len(frankenstein_freq))

def generate_frequency_list(path, limit = 0):
    with open(path) as f:
        file_contents = f.read()
        print (f"--- Begin report of books/frankenstein.txt ---")
        print ("\n")
        frequency = count_unique_words(file_contents)
        word_count = frequency[0][1] # This is always the total_words as it is the largest
        print (f"{word_count} words found in the document")
        pPrint(frequency, limit)
        print ("--- End report ---")
    return (word_count, frequency)


def count_unique_words(string):
    frequency = {}
    frequency["total_words"] = 0
    words = string.split()
    for word in words:
        word = word.lower()
        word = re.sub(r'[^\w\s]','',word)
        if word not in frequency:
            frequency[word] = 1
        else:
            frequency[word] += 1
        frequency["total_words"] += 1
    return sorted(frequency.items(), key=lambda x: x[1], reverse=True)


def pPrint(dict, limit = 0):
    if limit == 0:
        for key, value in dict:
            if key != "total_words":
                print(f"{key}: {value}")
            else:
                continue
    if limit != 0:
        line = 0
        for key, value in dict:
            if key == "total_words":
                continue
            if line >= limit:
                break
            print (f"{key}: {value}")
            line += 1
main()
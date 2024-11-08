import pandas as pd
import numpy as np

def load_corpus () -> list:
    corpus = ['data science is one of the most important fields of science',
            'this is one of the best data science courses',
            'data scientists analyze data' ]
    return corpus

def clean_file(string) -> list:
    # Pulled out because this just a simple clean
    words = string.split(' ')
    return words

def create_set(corpus) -> list:
    # Just a file with all words in it
    words_set = set()
    for doc in corpus:
        words = clean_file(doc)
        words_set = words_set.union(set(words))
    n_docs = len(corpus)
    n_words_set = len(words_set)
    return words_set, n_docs, n_words_set
        
def dt_tf_gen(corpus) -> pd.DataFrame:
    words_set, n_docs, n_words_set = create_set(corpus)
    df_tf = pd.DataFrame(np.zeros((n_docs, n_words_set)), columns=list(words_set))
    # This is an array of shape len(words_set), len
    for i in range (n_docs):
        words = corpus[i].split(' ')
        for w in words:
            df_tf.loc[i,w] = df_tf[w][i] + (1 / len(words)) # Sum of each version + 1/len(words)
    return df_tf

def idf_gen(corpus) ->pd.DataFrame:
    # This is just a single array, it works because all of the keys are the same
    words_set, n_docs, n_words_set = create_set(corpus)

    idf = {}

    for w in words_set:
        k = 0    # number of documents in the corpus that contain this word
        for i in range(n_docs):
            if w in corpus[i].split(' '):
                k += 1
                
        idf[w] =  np.log10(n_docs / k)
        
    return idf

def df_tf_idf(corpus):
    df_tf = dt_tf_gen(corpus)
    idf = idf_gen(corpus)
    df_tf_idf = df_tf.copy()
    words_set, n_docs, n_words_set = create_set(corpus)
    for w in words_set:
        for i in range(n_docs):
            df_tf_idf[i, w] = df_tf[w][i] * idf[w]
    return df_tf_idf

corpus = load_corpus()
print(df_tf_idf(corpus))
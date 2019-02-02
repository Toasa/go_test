def gen_location_dict(input):
    pos = 0
    locate_dict = {}
    bytes_seq = []
    dict_val = 1

    while 1:
        pos_from = pos
        pos_to = pos + 1
        while input[pos_from:pos_to] in locate_dict.keys():
            pos_to += 1
        new_bytes = input[pos_from:pos_to]
        bytes_seq.append(new_bytes)
        locate_dict[new_bytes] = format(dict_val, '03b')
        dict_val += 1
        pos = pos_to
        if pos >= len(input):
            break
    return locate_dict, bytes_seq

def encode(locate_dict, bytes_seq):
    code = ''
    for word in bytes_seq:
        if len(word) <= 1:
            code += '000'
        else:
            code += locate_dict[word[:-1]]
        code += word[-1]
        code += ' '
    
    return code

def main():
    input = '101011011010101010'
    d, l = gen_location_dict(input)
    # print(d)
    # print(l)
    print(encode(d, l))
    

if __name__ == '__main__':
    main()

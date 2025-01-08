/*
You are given an array of strings of the same length words.

In one move, you can swap any two even indexed characters or any two odd indexed characters of a string words[i].

Two strings words[i] and words[j] are special-equivalent if after any number of moves, words[i] == words[j].

For example, words[i] = "zzxy" and words[j] = "xyzz" are special-equivalent because we may make the moves "zzxy" -> "xzzy" -> "xyzz".
A group of special-equivalent strings from words is a non-empty subset of words such that:

Every pair of strings in the group are special equivalent, and
The group is the largest size possible (i.e., there is not a string words[i] not in the group such that words[i] is special-equivalent to every string in the group).
Return the number of groups of special-equivalent strings from words.
*/

/*
 * Analyse:
 *
 * For a word to be a special-equivalent it need to have the same number of letter in the odd and even position as the other words.
 *
 * We could do a first pass where we check the arrangement and another one to try to match them.
*/


function numSpecialEquivGroups(words: string[]): number {
  const wordsSet: Set<string> = new Set<string>();

  words.forEach((word) => {
    let oddChars: string[] = []
    let evenChars: string[] = []

    for (let index = 0; index < word.length; index++) {
      const element: string = word[index];
      index % 2 === 0 ? evenChars.push(element) : oddChars.push(element)
    }

    const totalChar = evenChars.sort().join('') + '_' + oddChars.sort().join('');

    wordsSet.add(totalChar);
  })

  return wordsSet.size;
}


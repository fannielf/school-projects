
const crosswordSolver = (emptyPuzzle, words) => {

    // Check for invalid string or array
    if (typeof emptyPuzzle !== 'string' || !Array.isArray(words) || !words.every(item => typeof item === "string")) {
        console.log('Error');
        return;
    }
  
    let validity = true;

    // Check reoccuring words
    for (let i=0; i<words.length; i++) {
        for (let j=i+1; j<words.length; j++) {
            if (words[i] === words[j]) {
                validity = false;
                break;
            }
        }
    }
  
    if (!validity) {
      console.log('Error');
      return
    }
  
    // Check string characters and empty string
    if (!/^[.\n012]+$/.test(emptyPuzzle)) {
        console.log('Error');
        return;
    }

    const numbers = emptyPuzzle.match(/\d/g);
    // Check that the amount of starting points equals the amount of words
    if (numbers === null || numbers.reduce((acc, num) => acc + parseInt(num), 0) !== words.length) {
        console.log('Error')
        return
    }
  
    // Sort 'words' from longest to shortest
    words.sort((a, b) => b.length - a.length);
    const shortest = words[words.length-1].length;
    const startingPoints = [];
    
    // Convert string into 2D array
    const grid = emptyPuzzle.split('\n').map(line => line.split(''));
    const height = grid.length;
    const width = grid[0].length;
  
    // Identify starting points for words (indexes with 1 or 2)
    for (let r = 0; r < height; r++) { // Iterates over rows
        for (let c = 0; c < width; c++) { // Iterates over columns
            if (grid[r][c] === '1' || grid[r][c] === '2') {
                // Check horizontally
                if (c === 0 || grid[r][c-1] === '0' || grid[r][c-1] === '.') {
                    let i = c;
                    while (i < width && (grid[r][i] === '0' || grid[r][i] === '1' || grid[r][i] === '2')) {
                        i++;
                    }
                    if (i - c >= shortest) { 
                        startingPoints.push({ row: r, col: c, direction: 'h' });
                    }
                }
                // Check vertically
                if (r === 0 || grid[r-1][c] === '0' || grid[r-1][c] === '.') {
                    let j = r;
                    while (j < height && (grid[j][c] === '0' || grid[j][c] === '1' || grid[j][c] === '2')) {
                        j++;
                    }
                    if (j - r >= shortest) { 
                        startingPoints.push({ row: r, col: c, direction: 'v' });
                    }
                }
            }
        }
    }
    let solutions = [];

    // Recursively solves the crossword puzzle
    const solve = (index) => {
  
      // Exit function if all words have been placed
        if (index === words.length) {
            solutions.push(grid.map(row => row.join('')).join('\n'));
            return;
        }
  
  
        // Try placing each word on grid
        const word = words[index];
        for (const start of startingPoints) {
            if (wordFits(word, start)) {
                placeWord(word, start);
                solve(index + 1)
                removeWord(word, start);
            }
        }
  
        return false;
    }
  
    // Check if a word can be placed at the given start position
    const wordFits = (word, start) => {
        let { row, col, direction } = start;
        for (let i = 0; i < word.length; i++) {
            if (row >= height || col >= width) return false;
            if (grid[row][col] !== '0' && grid[row][col] !== '1' && grid[row][col] !== '2' && grid[row][col] !== word[i]) {
                return false;
            }
  
            // Check if the end of the word and the next place is 0.
            // Proceed to next cell depending on direction
            if (direction === 'h') {
                if (i === word.length -1 && col+1 < width && grid[row][col + 1] === '0') {
                    return false;
                }
              col++
            } else {
                if (i === word.length -1 && row+1 < height && grid[row + 1][col] === '0') {
                    return false;
                }
              row++
            }
  
        }

        return true;
    }
  
    // Place the word at the start position
    const placeWord = (word, start) => {
        let { row, col, direction } = start;
        for (let i = 0; i < word.length; i++) {
            grid[row][col] = word[i];
            direction === 'h' ? col++ : row++;
        }
    }
  
    // Remove the word from the puzzle (used for backtracking)
    const removeWord = (word, start) => {
    
        let { row, col, direction } = start;
        for (let i = 0; i < word.length; i++) {
            grid[row][col] = '0';
  
              // Proceed to next cell depending on direction
              if (direction === 'h') {
                col++
              } else {
                row++
              }
        }
    }

    // Checks if the solution still has unused numbers
    const gridHasNumbers = (gridToCheck) => {
        for (let r = 0; r < gridToCheck.length; r++) {
            for (let c = 0; c < gridToCheck[0].length; c++) {
                if (gridToCheck[r][c] === '0' || gridToCheck[r][c] === '1' || gridToCheck[r][c] === '2') {
                    return true;
                }
            }
        }
        return false;
    }

    solve(0);

    const validSolutions = [];

    // Filter only valid solutions
    for (let i = 0; i < solutions.length; i++) {
        const solution = solutions[i];
        const solutionGrid = solution.split('\n').map(row => row.split(''));
        const hasNumbers = gridHasNumbers(solutionGrid);

        if (!hasNumbers) {
            validSolutions.push(solution);
        }
    }

    // Attempt to solve the puzzle and output result
    if (validSolutions.length === 1) {
        console.log(validSolutions[0]);
    } else {
        console.log('Error');
    }
  }
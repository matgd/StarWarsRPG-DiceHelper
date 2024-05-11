# StarWarsRPG-DiceHelper

TODO:

- [ ] Search bar
- [ ] Show chances of getting result (maybe dropdown?)
- [X] Add modifier
- [X] Focus for non-core attributes
- [X] Reload from JSON
- [X] Save to JSON
- [X] Show what to write in discord bot


## Show chances of getting result 
You can base tests on this:
https://www.gigacalculator.com/calculators/dice-probability-calculator.php
At least one die >= X

Example:

At least one 6 from 3 dices  
1 - (5/6 * 5/6 * 5/6) = 1 - 125/216 = 91/216 ~ 42%  

At least one 5 from 4 dices  
1 - (4/6 * 4/6 * 4/6 * 4/6) = 1 - 8/27 = 19/27 ~ 70%  

TODO: Explore more with at least N from M dices  

At least 2 6s from 3 dices  

1 minus  
0 6s (5/6 * 5/6 * 5/6) = 125/216  

1 6  (1/6 * 5/6 * 5/6) = 25/216  
1 6  (5/6 * 1/6 * 5/6) = 25/216  
1 6  (5/6 * 5/6 * 1/6) = 25/216  

1 - 200/216 = 16/216 ~ 7.4%  


At least 2 6s from 4 dices
1 minus
0 6s (5/6 * 5/6 * 5/6 * 5/6) = 625/1296

1 6 (1/6 * 5/6 * 5/6 * 5/6) = 125/1296
1 6 (5/6 * 1/6 * 5/6 * 5/6) = 125/1296
1 6 (5/6 * 5/6 * 1/6 * 5/6) = 125/1296
1 6 (5/6 * 5/6 * 5/6 * 1/6) = 125/1296


1 - 1125/1296 = 171/1296 ~ 13.2%  

At least 2 5s from 4 dices  
1 minus  
0 5s (4/6 * 4/6 * 4/6 * 4/6) = 256/1296  

1 5 (2/6 * 4/6 * 4/6 * 4/6) = 128/1296  
1 5 (4/6 * 2/6 * 4/6 * 4/6) = 128/1296  
1 5 (4/6 * 4/6 * 2/6 * 4/6) = 128/1296  
1 5 (4/6 * 4/6 * 4/6 * 2/6) = 128/1296  

1 - 640/1296 = 656/1296 ~ 50.6%  

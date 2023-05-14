const defaultCodeSnippets = {
    cpp: `class Solution {
            public:
                void/int/flote functionName(parameters) {
                    
                }
            }; \n`,
    javascript: `let solution = function(parameters) {

                }; \n`,
    java: `class Solution {
            public void/int/float function(parameters) {
                
            }
        }\n`,
    go: `func functionName(parameters) int/void/float {

    }\n`,
    python: `class Solution:
        def functionName(parameters) :
    \n`,
}

export default defaultCodeSnippets;
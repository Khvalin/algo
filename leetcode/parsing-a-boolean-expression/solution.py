class Solution:
    def parseBoolExpr(self, expression: str) -> bool:
        stack = []

        def evaluate():
            operator = None
            operands = []
            f = True
            while len(stack) and f:
                ch = stack.pop()
                if isinstance(ch, bool):
                    operands.append(ch)
                    continue
                if ch == "f" or ch == "t":
                    operands.append(ch == "t")
                    continue
                operator = ch
                f = False
            if operator is None:
                return bool(len(operands) and operands[0])
            if operator == "!":
                return not operands[0]
            if operator == "&":
                return all(operands)
            if operator == "|":
                return any(operands)

        # expression = f"|({expression})"
        for ch in expression:
            if ch == "(" or ch == ",":
                continue
            if ch != ")":
                stack.append(ch)
                continue

            res = evaluate()
            stack.append(res)

        return evaluate()

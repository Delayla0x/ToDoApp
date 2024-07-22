import React, { useMemo } from 'react';

function ExpensiveComponent({ value }) {
    const expensiveCalculation = useMemo(() => {
        return calculateExpensiveValue(value);
    }, [value]); 

    return <div>{expensiveCalculation}</div>;
}
```
```tsx
import React, { useCallback } from 'react';

function ParentComponent(props) {
    const memoizedCallback = useCallback(
        () => {
        },
        [], 
    );

    return <ChildComponent onSomeEvent={memoizedCallback} />;
}
# ARI

Golang Implementation of the JavaScript [Array](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array) Api. With Support for the following methods:
## ArrayList
|Method|Description|Parameters|Return|
|---|---|---|---|
|`Add()`|Adds an element to the end of an array|`any`|N/A|
|`Clear()`|Removes all elements from an array|N/A|N/A|
|`Copy()`|Creates a shallow copy of an array list|N/A|`*ArrayList`|
|`Every()`|Tests whether all elements in the array pass the test implemented by the provided function|`func(any) bool`|`bool`|
|`Filter()`|Creates a new array with all elements that pass the test implemented by the provided function|`func(any) bool`|`*ArrayList`|
|`Find()`|Returns the value of the first element in the array that satisfies the provided testing function|`func(any) bool`|`any`|
|`FindIndex()`|Returns the index of the first element in the array that satisfies the provided testing function|`func(any) bool`|`int`|
|`ForEach()`|Calls a function for each element in the array <u>**sequentially**</u> |`func(any)`|N/A|
|`ForEachAsync()`|Calls a function for each element in the array <u>**concurrently**</u> *|`func(any)`|N/A|
|`Get()`|Returns the element at the specified index in the array|`int`|`any`|
|`Includes()`|Determines whether an array includes a certain value among its entries, returning true or false as appropriate|`any`|`bool`|
|`IsEmpty()`|Determines whether an array is empty|N/A|`bool`|
|`IsNotEmpty()`|Determines whether an array is not empty|N/A|`bool`|
|`Map()`| returns a new list contianing the results of applying the given function|`func(any) any`|`*ArrayList`|
|`MapBy()`| returns a hash map of the list using the provided string as key|`func(func(any)(string,any())map[string]any)`|map[string]any|
|`Peek()`|Returns the last element in the array|N/A|`any`|
|`PeekFirst()`|Returns the first element in the array|N/A|`any`|
|`PeekLast()`|Returns the last element in the array|N/A|`any`|
|`PeekNthFromLast()`|Returns the nth element from the end of the array|`int`|`any`|
|`PeekRandom()`|Returns a random element in the array|N/A|`any`|
|`ReduceRight()`|Applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value|`func(any, any) any`|`any`|
|`Reduce()`|Applies a function against an accumulator and each value of the array (from left-to-right) to reduce it to a single value|`func(any, any) any`|`any`|
|`Remove()`|Removes the first occurrence of a value from an array and returns it|`any`|`any`|
|`Set()`|Sets the element at the specified index in the array|`int`, `any`|N/A|
|`Shift()`|Removes the first element from an array and returns it|N/A|`any`|
|`Slice()`|Returns a shallow copy of a portion of an array into a new array object|`int`, `int`|`*ArrayList`|
|`Some()`|Tests whether at least one element in the array passes the test implemented by the provided function|`func(any) bool`|`bool`|
|`String()`|Returns a string representation of the array|N/A|`string`|
|`ToArray()`|Returns a new array containing a shallow copy of the elements in the array|N/A|`[]any`|
|`ToChainedMapString()`|Returns a new chained map containing a shallow copy of the elements in the array|N/A|`map[string][]any`|
|`ToMapString()`|Returns a new map containing a shallow copy of the elements in the array|N/A|`map[string]any`|
|`ToSet()`|Returns a new set containing a shallow copy of <u>**UNIQUE**</u> elements in the array|N/A|`*Set`|
|`ToAlice()`|Returns a new array containing a shallow copy of the elements in the array|N/A|`[]any`|
|`Unshift()`|Adds one or more elements to the beginning of an array and returns the new length of the array|`any`|`int`|

## Constructor Methods
|Method|Description|Parameters|Return|
|---|---|---|---|
|`InitFromSlice()`|Creates a new ArrayList from a slice|`[]any`|`*ArrayList`|
|`InitFromHashMap()`|Creates a new ArrayList from a HashMap|`map[string]any`|`*ArrayList`|

## Tuple
|Method|Description|Parameters|Return|
|---|---|---|---|
|`ToString()`|Returns a string representation of the tuple|N/A|`string`|
|`Key()`|Returns the key of the tuple|N/A|`any`|
|`Value()`|Returns the value of the tuple|N/A|`any`|
|`SetKey()`|Sets the key of the tuple|`any`|N/A|
|`SetValue()`|Sets the value of the tuple|`any`|N/A|
|`Equals()`|Determines whether two tuples are equal via string comparison|`*Tuple`|`bool`|
|`EqualsByComparator()`|Determines whether two tuples are equal via a comparator function|`*Tuple`, `func(any, any) bool`|`bool`|




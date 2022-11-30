# ARI

Golang Implementation of the JavaScript [Array](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array) Api. With Support for the following methods:
|Method|Description|Parameters|Return|
|---|---|---|---|
|`Add()`|Adds an element to the end of an array|`any`|N/A|
|`Clear()`|Removes all elements from an array|N/A|N/A|
|`Copy()`|Creates a shallow copy of an array list|N/A|`*ArrayList`|
|`Every()`|Tests whether all elements in the array pass the test implemented by the provided function|`func(interface{}) bool`|`bool`|
|`Filter()`|Creates a new array with all elements that pass the test implemented by the provided function|`func(interface{}) bool`|`*ArrayList`|
|`Find()`|Returns the value of the first element in the array that satisfies the provided testing function|`func(interface{}) bool`|`interface{}`|
|`FindIndex()`|Returns the index of the first element in the array that satisfies the provided testing function|`func(interface{}) bool`|`int`|
|`ForEach()`|Calls a function for each element in the array <u>**sequentially**</u> |`func(interface{})`|N/A|
|`ForEachAsync()`|Calls a function for each element in the array <u>**concurrently**</u> *|`func(interface{})`|N/A|
|`Get()`|Returns the element at the specified index in the array|`int`|`interface{}`|
|`Includes()`|Determines whether an array includes a certain value among its entries, returning true or false as appropriate|`any`|`bool`|
|`IsEmpty()`|Determines whether an array is empty|N/A|`bool`|
|`IsNotEmpty()`|Determines whether an array is not empty|N/A|`bool`|
|`Map()`| returns a new list contianing the results of applying the given function|`func(interface{}) interface{}`|`*ArrayList`|
|`MapBy()`| returns a hash map of the list using the provided string as key|`func(func(interface{})(string,interface())map[string]interface)`|map[string]interface{}|
|`Peek()`|Returns the last element in the array|N/A|`interface{}`|
|`PeekFirst()`|Returns the first element in the array|N/A|`interface{}`|
|`PeekLast()`|Returns the last element in the array|N/A|`interface{}`|
|`PeekNthFromLast()`|Returns the nth element from the end of the array|`int`|`interface{}`|
|`PeekRandom()`|Returns a random element in the array|N/A|`interface{}`|
|`ReduceRight()`|Applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value|`func(interface{}, interface{}) interface{}`|`interface{}`|
|`Reduce()`|Applies a function against an accumulator and each value of the array (from left-to-right) to reduce it to a single value|`func(interface{}, interface{}) interface{}`|`interface{}`|
|`Remove()`|Removes the first occurrence of a value from an array and returns it|`any`|`interface{}`|
|`Set()`|Sets the element at the specified index in the array|`int`, `any`|N/A|
|`Shift()`|Removes the first element from an array and returns it|N/A|`interface{}`|
|`Slice()`|Returns a shallow copy of a portion of an array into a new array object|`int`, `int`|`*ArrayList`|
|`Some()`|Tests whether at least one element in the array passes the test implemented by the provided function|`func(interface{}) bool`|`bool`|
|`String()`|Returns a string representation of the array|N/A|`string`|
|`ToArray()`|Returns a new array containing a shallow copy of the elements in the array|N/A|`[]interface{}`|
|`ToSet()`|Returns a new set containing a shallow copy of <u>**UNIQUE**</u> elements in the array|N/A|`*Set`|
|`ToAlice()`|Returns a new array containing a shallow copy of the elements in the array|N/A|`[]interface{}`|
|`Unshift()`|Adds one or more elements to the beginning of an array and returns the new length of the array|`any`|`int`|



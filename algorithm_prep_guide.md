# 4-Day Algorithm Interview Preparation Plan

## День 1: Основы и базовые структуры данных / Basics & Data Structures

### 1.1 Arrays & Strings (Массивы и строки)

**Что повторить:**
- Итерация, индексы, срезы
- Two pointers technique
- Sliding window
- In-place operations

**Примеры задач:**
1. **Two Sum** — найти два числа, дающих целевую сумму
2. **Reverse String** — развернуть строку in-place
3. **Valid Anagram** — проверить, являются ли строки анаграммами
4. **Longest Substring Without Repeating Characters** — найти самую длинную подстроку без повторяющихся символов
5. **Container With Most Water** — найти контейнер с максимальным объёмом

### 1.2 Hash Maps & Sets (Хеш-мапы и множества)

**Что повторить:**
- Когда использовать hash map vs array
- Обработка коллизий (теория)
- Подсчёт частот элементов

**Примеры задач:**
1. **Two Sum** (решение через hash map за O(n))
2. **Group Anagrams** — сгруппировать анаграммы
3. **Top K Frequent Elements** — найти K самых частых элементов
4. **Longest Consecutive Sequence** — найти самую длинную последовательность подряд

### 1.3 Linked Lists (Связные списки)

**Что повторить:**
- Singly vs Doubly linked list
- Fast & slow pointers (Floyd's cycle detection)
- Reversing a linked list

**Примеры задач:**
1. **Reverse Linked List** — развернуть связный список
2. **Middle of the Linked List** — найти середину списка
3. **Linked List Cycle** — определить, есть ли цикл
4. **Merge Two Sorted Lists** — слить два отсортированных списка

---

## День 2: Рекурсия, деревья и графы / Recursion, Trees & Graphs

### 2.1 Recursion (Рекурсия)

**Что повторить:**
- Base case и recursive case
- Стек вызовов
- Memoization

**Примеры задач:**
1. **Fibonacci Number** (с memoization)
2. **Climbing Stairs** — сколькими способами подняться по лестнице
3. **Power of Three** — является ли число степенью тройки

### 2.2 Binary Trees & BST (Бинарные деревья и BST)

**Что повторить:**
- DFS: Pre-order, In-order, Post-order
- BFS (level-order traversal)
- Свойства BST
- Высота/глубина дерева

**Примеры задач:**
1. **Maximum Depth of Binary Tree** — найти максимальную глубину
2. **Invert Binary Tree** — инвертировать дерево
3. **Validate Binary Search Tree** — проверить, является ли дерево BST
4. **Level Order Traversal** — обход дерева по уровням
5. **Lowest Common Ancestor of BST** — найти наименьшего общего предка

### 2.3 Graphs (Графы)

**Что повторить:**
- DFS и BFS на графах
- Представление графов: adjacency list vs matrix
- Поиск компонент связности
- Топологическая сортировка (базово)

**Примеры задач:**
1. **Number of Islands** — посчитать количество островов
2. **Clone Graph** — клонировать граф
3. **Course Schedule** (топологическая сортировка)
4. **Find if Path Exists in Graph** — проверить существование пути

---

## День 3: Продвинутые техники / Advanced Techniques

### 3.1 Binary Search (Бинарный поиск)

**Что повторить:**
- Классический бинарный поиск
- Поиск в rotated sorted array
- Поиск ответа (binary search on answer)

**Примеры задач:**
1. **Binary Search** — базовая реализация
2. **Search in Rotated Sorted Array**
3. **Find First and Last Position of Element**
4. **Find Minimum in Rotated Sorted Array**
5. **Koko Eating Bananas** (binary search on answer)

### 3.2 Sliding Window & Two Pointers (Скользящее окно и два указателя)

**Что повторить:**
- Fixed-size vs variable-size window
- Когда использовать sliding window

**Примеры задач:**
1. **Best Time to Buy and Sell Stock**
2. **Longest Repeating Character Replacement**
3. **Minimum Window Substring**
4. **3Sum** — найти три числа с суммой 0

### 3.3 Backtracking (Возврат)

**Что повторить:**
- Шаблон backtracking: choose → explore → un-choose
- Pruning (отсечение ветвей)

**Примеры задач:**
1. **Permutations** — все перестановки
2. **Subsets** — все подмножества
3. **Combination Sum** — комбинации с заданной суммой
4. **Word Search** — найти слово в матрице
5. **Generate Parentheses** — сгенерировать корректные скобки

---

## День 4: Динамическое программирование и жадные алгоритмы / DP & Greedy

### 4.1 Dynamic Programming (Динамическое программирование)

**Что повторить:**
- Top-down (memoization) vs Bottom-up (tabulation)
- Как определить подзадачу
- Одномерная и двумерная DP

**Примеры задач:**
1. **Climbing Stairs** (одномерная DP)
2. **House Robber** (одномерная DP)
3. **Coin Change** — минимальное количество монет
4. **Longest Increasing Subsequence** — самая длинная возрастающая подпоследовательность
5. **Longest Common Subsequence** — самая длинная общая подпоследовательность
6. **Unique Paths** — количество уникальных путей в матрице

### 4.2 Greedy Algorithms (Жадные алгоритмы)

**Что повторить:**
- Когда жадный подход работает
- Как доказать корректность (интуиция)

**Примеры задач:**
1. **Maximum Subarray** (алгоритм Кадане)
2. **Jump Game** — можно ли допрыгать до конца
3. **Non-overlapping Intervals** — максимальное количество непересекающихся интервалов

### 4.3 Intervals & Heaps (Интервалы и кучи)

**Что повторить:**
- Сортировка интервалов
- Min-heap и max-heap
- K-way merge

**Примеры задач:**
1. **Merge Intervals** — слить перекрывающиеся интервалы
2. **Kth Largest Element in an Array** — K-й наибольший элемент
3. **Top K Frequent Elements** (через heap)
4. **Meeting Rooms II** — минимальное количество комнат

---

## Чек-лист перед интервью / Pre-Interview Checklist

- [ ] Знаю Big O для основных операций (array, hash map, BST)
- [ ] Могу реализовать binary search без багов
- [ ] Понимаю DFS/BFS и могу применить к деревьям и графам
- [ ] Могу распознать DP задачу и написать рекуррентное соотношение
- [ ] Знаю шаблон backtracking
- [ ] Понимаю when to use sliding window
- [ ] Могу развернуть linked list
- [ ] Знаю, как найти цикл в графе

## Полезные ресурсы / Useful Resources

- **LeetCode** — практика задач (Blind 75 / NeetCode 150)
- **NeetCode.io** — структурированные задачи по темам
- **AlgoExpert** — видео-объяснения
- **Grokking the Coding Interview** — курс по паттернам

## Советы по подготовке / Preparation Tips

1. **Практикуйся вслух** — проговаривай решение, как на интервью
2. **Не зубри** — понимай паттерны, а не конкретные решения
3. **Тайминг** — решай задачу за 20-30 минут
4. **Разбор ошибок** — если не решил, разбери решение и реши заново через день
5. **Mock interviews** — потренируйся с другом или на Pramp/Interviewing.io

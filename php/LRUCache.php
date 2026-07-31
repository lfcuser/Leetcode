<?php

/*
    146 LRU Cache

    Implement the LRUCache class:

    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
    int get(int key) Return the value of the key if the key exists, otherwise return -1.
    void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
    If the number of keys exceeds the capacity from this operation, evict the least recently used key.

    The functions get and put must each run in O(1) average time complexity.
*/

class Node
{
    public function __construct(
        public $key = null,
        public $value = null,
        public ?Node $next = null,
        public ?Node $prev = null,
    ) {}
}

class LRUCache {
    public array $map = [];
    public ?Node $head = null;
    public ?Node $tail = null;

    /**
     * @param Integer $capacity
     */
    function __construct(public int $capacity)
    {
        $this->head = new Node();
        $this->tail = new Node();

        $this->tail->next = $this->head;
        $this->head->prev = $this->tail;
    }

    private function remove(Node $node): void
    {
        $node->prev->next = $node->next;
        $node->next->prev = $node->prev;
    }

    private function addToMRU(Node $node): void
    {
        $prevNode = $this->head->prev;

        $prevNode->next = $node;
        $node->prev = $prevNode;

        $node->next = $this->head;
        $this->head->prev = $node;
    }

    /**
     * @param Integer $key
     * @return Integer
     */
    function get($key)
    {
        if (!isset($this->map[$key])) {
            return -1;
        }

        $node = $this->map[$key];

        $this->remove($node);
        $this->addToMRU($node);

        return $node->value;
    }

    /**
     * @param Integer$key
     * @param Integer $value
     * @return NULL
     */
    function put($key, $value)
    {
        if (isset($this->map[$key])) {
            $node = $this->map[$key];
            $node->value = $value;

            $this->remove($node);
            $this->addToMRU($node);
            return;
        }

        if (count($this->map) === $this->capacity) {
            $lruNode = $this->tail->next;

            $this->remove($lruNode);
            unset($this->map[$lruNode->key]);
        }

        $newNode = new Node(key: $key, value: $value);
        $this->addToMRU($newNode);
        $this->map[$key] = $newNode;
    }
}

$obj = new LRUCache(1);
$ret_1 = $obj->get(1);
var_dump($ret_1);
$obj->put(1, 1);
var_dump($obj->map);
$obj->put(2, 2);
var_dump($obj->map);

// php ./php/LRUCache.php

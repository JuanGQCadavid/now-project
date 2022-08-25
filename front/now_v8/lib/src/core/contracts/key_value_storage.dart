

abstract class IKeyValueStorage<K,V> {
  V getValue(K key);
  save(V value, K key);
  update(V value, K key);
}
import 'package:dartz/dartz.dart';

abstract class IKeyValueStorage<K, V> {
  Either<V, None> getValue(K key);
  Future delete(K key);
  save(V value, K key);
  update(V value, K key);
  Future doInit();
}

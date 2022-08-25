


import 'package:hive/hive.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';

class HiveKeyValue<V> implements IKeyValueStorage<String,V>{

  late Box box;

  HiveKeyValue(String boxName){
    box = Hive.box(boxName);
  }

  @override
  V getValue(String key) {
    return box.get(key);
  }

  @override
  save(V value, String key) {
    box.put(key, value);
  }

  @override
  update(V value,String key) {
    box.put(key, value);
  }

}
import 'package:hive/hive.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';

class HiveKeyValue<V> implements IKeyValueStorage<String, V> {
  final String boxName;
  HiveKeyValue({required this.boxName});

  doInit() async{
    if(!Hive.isBoxOpen(boxName)) {
      await Hive.openBox(boxName);
    }
  }

  @override
  V getValue(String key) {
    var box = Hive.box(boxName);
    return box.get(key, defaultValue: "");
  }

  @override
  save(V value, String key) {
    var box = Hive.box(boxName);
    box.put(key, value);
  }

  @override
  update(V value, String key) {
    var box = Hive.box(boxName);
    box.put(key, value);
  }
}

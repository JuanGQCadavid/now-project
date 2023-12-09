import 'package:dartz/dartz.dart';
import 'package:hive/hive.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';

class HiveKeyValue<V> implements IKeyValueStorage<String, V> {
  final String boxName;
  HiveKeyValue({required this.boxName}) {
    print("Initializing myself");
    doInit();
  }

  doInit() async {
    if (!Hive.isBoxOpen(boxName)) {
      await Hive.openBox(boxName);
      print("I'am initialized");
    }
  }

  @override
  Either<V, None> getValue(String key) {
    var box = Hive.box(boxName);
    var value = box.get(key);

    if (value == null) {
      return right(const None());
    }

    return left(value);
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

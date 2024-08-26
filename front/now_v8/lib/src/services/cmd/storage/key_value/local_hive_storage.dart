import 'package:dartz/dartz.dart';
import 'package:hive/hive.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';

class HiveKeyValue<V> implements IKeyValueStorage<String, V> {
  final String boxName;
  HiveKeyValue({required this.boxName}) {
    doInit();
  }

  @override
  doInit() async {
    if (!Hive.isBoxOpen(boxName)) {
      await Hive.openBox<V>(boxName);
    }
  }

  @override
  Either<V, None> getValue(String key) {
    var box = Hive.box<V>(boxName);
    var value = box.get(key);

    if (value == null) {
      return right(const None());
    }

    return left(value);
  }

  @override
  save(V value, String key) {
    var box = Hive.box<V>(boxName);
    box.put(key, value);
  }

  @override
  update(V value, String key) {
    var box = Hive.box<V>(boxName);
    box.put(key, value);
  }

  @override
  Future delete(String key) async {
    var box = Hive.box<V>(boxName);
    await box.delete(key);
  }
}

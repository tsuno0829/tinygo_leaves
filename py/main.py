from sklearn import datasets
from sklearn.model_selection import train_test_split
import lightgbm as lgb


def main():
    params = {
        'objective': 'binary',
        'metric': 'binary_logloss',
        'learning_rate': 0.1,
        'n_estimators': 100,
        'verbose': -1,
    }

    iris = datasets.load_breast_cancer()
    X, y = iris.data, iris.target
    X_train, X_test, y_train, y_test = train_test_split(X, y, random_state=42)

    lgb_train = lgb.Dataset(X_train, y_train)
    model = lgb.train(params, lgb_train, valid_sets=[lgb_train])

    y_pred = model.predict(X_test)
    model.save_model('./go/model.txt')
    print(X_test[:1])
    print(y_test[:1])
    print(y_pred[:1])


if __name__ == '__main__':
    main()

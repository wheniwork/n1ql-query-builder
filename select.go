package nqb

func Select(expressions ...*Expression) FromPath {
	return newDefaultSelectPath(nil).SelectExpr(expressions...)
}

/*
public class Select {

    private Select() {}

    public static FromPath select(Expression... expressions) {
        return new DefaultSelectPath(null).select(expressions);
    }

    public static FromPath select(String... expressions) {
        return new DefaultSelectPath(null).select(expressions);
    }

    public static FromPath selectAll(Expression... expressions) {
        return new DefaultSelectPath(null).selectAll(expressions);
    }

    public static FromPath selectAll(String... expressions) {
        return new DefaultSelectPath(null).selectAll(expressions);
    }

    public static FromPath selectDistinct(Expression... expressions) {
        return new DefaultSelectPath(null).selectDistinct(expressions);
    }

    public static FromPath selectDistinct(String... expressions) {
        return new DefaultSelectPath(null).selectDistinct(expressions);
    }

    public static FromPath selectRaw(Expression expression) {
        return new DefaultSelectPath(null).selectRaw(expression);
    }

    public static FromPath selectRaw(String expression) {
        return new DefaultSelectPath(null).selectRaw(expression);
    }
}

*/

